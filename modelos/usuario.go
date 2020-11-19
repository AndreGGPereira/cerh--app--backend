package modelos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/andreggpereira/cerh--app--backend/controler"
	"github.com/andreggpereira/cerh--app--backend/persistencia"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

//Usuario capitalize to export from package
type Usuario struct {
	ID             int    `json:"id,omitempty"`
	Nome           string `json:"nome,omitempty"`
	NomeGestor     string `json:"nomegestor,omitempty"`
	Telefone       string `json:"telefone,omitempty"`
	Login          string `json:"login,omitempty"`
	Senha          []byte `json:"senha,omitempty"`
	DataCadastro   string `json:"datacadastro,omitempty"`
	Auditoria      string `json:"auditoria,omitempty"`
	PrimeiroAcesso bool   `json:"primeiroacesso,omitempty"`
	Ativo          bool   `json:"ativo,omitempty"`
	DataAtivo      string `json:"dataativo,omitempty"`
	Validado       bool   `json:"validade,omitempty"`
	DataValidade   string `json:"datavalidade,omitempty"`

	Empresa Empresa `json:"empresa,omitempty"`
}

var usuarios []Usuario

func insertORUpUsuario(obj Usuario) int {

	var id int
	//Se exister ID item deverá ser atualizado
	if obj.ID != 0 {
		id = obj.ID
		stmt, err := persistencia.DB.Prepare("update Usuario set nome = $1, telefone = $2, login = $3, senha = $4, datacadastro = $5, id_empresa = $6 where id =$7")
		stmt.Exec(obj.Nome, obj.Telefone, obj.Login, obj.Senha, obj.DataCadastro, obj.Empresa.ID, obj.ID)
		if err != nil {
			//log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	} else {

		stmt, err := persistencia.DB.Prepare("insert into Usuario(nome, nomegestor, telefone, login, senha, datacadastro, auditoria, primeiroacesso, ativo, dataativo, validade, datavalidade, id_empresa) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)  RETURNING id")
		err = stmt.QueryRow(obj.Nome, obj.NomeGestor, obj.Telefone, obj.Login, obj.Senha, obj.DataCadastro, obj.Auditoria, obj.PrimeiroAcesso, obj.Ativo, obj.DataAtivo, obj.Validado, obj.DataValidade, obj.Empresa.ID).Scan(&id)
		fmt.Println("Teste ID ", &id, id)
		if err != nil {
			//log.Fatal("Cannot run insert statement", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	}
	return id
}

//UsuarioLogin login do usuario
func UsuarioLogin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var gestor Usuario
	var gestorDB Usuario
	err := json.NewDecoder(req.Body).Decode(&gestor)
	if err != nil {
		http.Error(w, "Ocorreu um erro ao realizar a operação", 403)
	}

	gestorDB = getUsuarioLogin(gestor.Login)

	if gestorDB.ID == 0 {
		json.NewEncoder(w).Encode(controler.Message{Message: "Login ou senha, inválidos", Status: 403})

	} else {
		//comparação do hash cadastrado no banco, com o hash da senha que foi digitada
		err := bcrypt.CompareHashAndPassword(gestorDB.Senha, []byte(gestor.Senha))

		if err != nil {
			json.NewEncoder(w).Encode(controler.Message{Message: "Login ou senha, inválidos", Status: 403})
		} else {

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"gestor":  gestorDB.ID,
				"empresa": gestorDB.Empresa.ID,
			})

			claims := token.Claims.(jwt.MapClaims)
			claims["autorized"] = true
			claims["user"] = gestorDB.Login
			claims["exp"] = time.Now().Add(time.Minute * 3000).Unix()

			tokenString, error := token.SignedString([]byte(controler.JwtSecretKey))

			if error != nil {
				fmt.Println("Erro token", error)
			}

			json.NewEncoder(w).Encode(controler.JwtToken{Token: tokenString, Status: 202, Login: gestorDB.Login, Nome: gestorDB.Nome})
		}
	}
}

func getUsuarioLogin(login string) Usuario {

	resultado, _ := persistencia.DB.Query("select id, nome, login, senha, id_empresa from gestor where login = $1", login)
	defer resultado.Close()
	var um Usuario
	for resultado.Next() {
		resultado.Scan(&um.ID, &um.Nome, &um.Login, &um.Senha, &um.Empresa.ID)
		um = Usuario{ID: um.ID, Nome: um.Nome, Login: um.Login, Senha: um.Senha, Empresa: um.Empresa}
	}

	return um
}

//CreateUsuario cadastro
func CreateUsuario(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var obj Usuario
	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.DataCadastro = controler.PegarDataAtualStringNew()

	obj.Empresa.ID, _ = controler.PegarIDEmpresaIDGestor(req)

	hashSenha, err := bcrypt.GenerateFromPassword([]byte(obj.Senha), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Erro ao gerar o Hash")
	}
	obj.Senha = hashSenha

	insertORUpUsuario(obj)
	json.NewEncoder(w).Encode("Gestor cadastrado com sucesso")
}

//GetUsuarioAll pegar todos os itens
func GetUsuarioAll(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var obj Usuario
	obj.Empresa.ID, _ = controler.PegarIDEmpresaIDGestor(req)

	usuarios = getUsuarioALL(obj.Empresa.ID)

	json.NewEncoder(w).Encode(&usuarios)
}
func getUsuarioALL(idempresa int) []Usuario {
	rows, err := persistencia.DB.Query("SELECT u.id, u.nome, u.datacadastro, u.login, u.telefone, p.nome FROM Usuario as u WHERE u.id_empresa =$1 ORDER BY u.nome ASC", idempresa)

	if err != nil {
		fmt.Println(err)
	}

	var ps []Usuario
	for rows.Next() {
		var obj Usuario

		rows.Scan(&obj.ID, &obj.Nome, &obj.DataCadastro, &obj.Login, &obj.Telefone)
		ps = append(ps, Usuario{ID: obj.ID, Nome: obj.Nome, DataCadastro: obj.DataCadastro, Login: obj.Login, Telefone: obj.Telefone})
	}

	defer rows.Close()
	return ps
}

//DeleteUsuario deleteFornecedor
func DeleteUsuario(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var obj Usuario
	var err error
	vars := mux.Vars(req)

	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Erro ao inserir o identificador")
	}

	deletarUsuario(obj.ID)

	//json.NewEncoder(w).Encode(Gestors)
}

//DeletarGestor deletar usuaria
func deletarUsuario(id int) {

	stmt2, err := persistencia.DB.Prepare("delete from Usuario where id = $1")
	stmt2.Exec(id)

	if err != nil {
		panic(err)
	}
	stmt2.Close()
}

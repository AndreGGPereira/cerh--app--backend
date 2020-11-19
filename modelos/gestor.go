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

//Gestor capitalize to export from package
type Gestor struct {
	ID             int     `json:"id,omitempty"`
	Nome           string  `json:"nome,omitempty"`
	NomeGestor     string  `json:"nomegestor,omitempty"`
	Telefone       string  `json:"telefone,omitempty"`
	Login          string  `json:"login,omitempty"`
	Senha          []byte  `json:"senha,omitempty"`
	DataCadastro   string  `json:"datacadastro,omitempty"`
	Auditoria      string  `json:"auditoria,omitempty"`
	PrimeiroAcesso bool    `json:"primeiroacesso,omitempty"`
	Ativo          bool    `json:"ativo,omitempty"`
	DataAtivo      string  `json:"dataativo,omitempty"`
	Validado       bool    `json:"validade,omitempty"`
	DataValidade   string  `json:"datavalidade,omitempty"`
	Empresa        Empresa `json:"empresa,omitempty"`
}

var gestores []Gestor

//GestorList struct consulta
type GestorList struct {
	Gestor   []Gestor          `json:"gestor,omitempty"`
	Contador int               `json:"contador,omitempty"`
	Message  controler.Message `json:"message,omitempty"`
}

func insertORUpGestor(obj Gestor) int {

	var id int
	//Se exister ID item deverá ser atualizado
	if obj.ID != 0 {
		id = obj.ID
		stmt, err := persistencia.DB.Prepare("update Gestor set nome = $1, nomegestor = $2, telefone = $3, login = $4, senha = $5, datacadastro = $6, auditoria = $7, primeiroacesso = $8, ativo = $9, dataativo = $10, validade = $11, datavalidade = $12, id_empresa = $13 where id =$14")
		err = stmt.QueryRow(obj.Nome, obj.NomeGestor, obj.Telefone, obj.Login, obj.Senha, obj.DataCadastro, obj.Auditoria, obj.PrimeiroAcesso, obj.Ativo, obj.DataAtivo, obj.Validado, obj.DataValidade, obj.Empresa.ID, obj.ID).Scan(&id)
		if err != nil {
			//log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	} else {

		stmt, err := persistencia.DB.Prepare("insert into gestor(nome, nomegestor, telefone, login, senha, datacadastro, auditoria, primeiroacesso, ativo, dataativo, validade, datavalidade, id_empresa) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)  RETURNING id")
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

//GestorLogin login do usuario
func GestorLogin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var gestor Gestor
	var gestorDB Gestor
	err := json.NewDecoder(req.Body).Decode(&gestor)
	if err != nil {
		http.Error(w, "Ocorreu um erro ao realizar a operação", 403)
	}

	gestorDB = getGestorLogin(gestor.Login)

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

func getGestorLogin(login string) Gestor {

	resultado, _ := persistencia.DB.Query("select id, nome, login, senha, id_empresa from gestor where login = $1", login)
	defer resultado.Close()
	var um Gestor
	for resultado.Next() {
		resultado.Scan(&um.ID, &um.Nome, &um.Login, &um.Senha, &um.Empresa.ID)
		um = Gestor{ID: um.ID, Nome: um.Nome, Login: um.Login, Senha: um.Senha, Empresa: um.Empresa}
	}

	return um
}

//CreateGestor cadastro
func CreateGestor(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var obj Gestor
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

	insertORUpGestor(obj)
	json.NewEncoder(w).Encode("Gestor cadastrado com sucesso")
}

//GetGestorAll pegar todos os itens
func GetGestorAll(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var obj Gestor
	var objList GestorList
	var err error
	obj.Empresa.ID, _ = controler.PegarIDEmpresaIDGestor(req)

	objList.Gestor, err = getGestorALL(obj.Empresa.ID)
	objList.Contador = len(objList.Gestor)

	if err != nil {
		objList.Message.Message = "Nao possivel realizar a consulta"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 200
	}

	json.NewEncoder(w).Encode(&objList)
}
func getGestorALL(idempresa int) ([]Gestor, error) {
	rows, err := persistencia.DB.Query("SELECT id, nome, nomegestor, telefone, login, datacadastro, auditoria, primeiroacesso, ativo, dataativo, validade, datavalidade FROM gestor WHERE id_empresa =$1", idempresa)

	if err != nil {
		fmt.Println(err)
	}

	var ps []Gestor
	for rows.Next() {
		var obj Gestor
		rows.Scan(&obj.ID, &obj.Nome, &obj.NomeGestor, &obj.Telefone, &obj.Login, &obj.DataCadastro, &obj.Auditoria, &obj.PrimeiroAcesso, &obj.Ativo, &obj.DataAtivo, &obj.Validado, &obj.DataValidade)
		ps = append(ps, Gestor{ID: obj.ID, Nome: obj.Nome, NomeGestor: obj.NomeGestor, Telefone: obj.Telefone, Login: obj.Login, DataCadastro: obj.DataCadastro, Auditoria: obj.Auditoria, PrimeiroAcesso: obj.PrimeiroAcesso, Ativo: obj.Ativo, DataAtivo: obj.DataAtivo, Validado: obj.Validado, DataValidade: obj.DataValidade})
	}

	defer rows.Close()
	return ps, err
}

// func getGestorALL(idempresa int) []Gestor {
// 	rows, err := persistencia.DB.Query("SELECT id, nome, nomegestor, telefone, login, senha, datacadastro, auditoria, primeiroacesso, ativo, dataativo, validade, datavalidade FROM gestor WHERE id_empresa =$1", idempresa)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var ps []Gestor
// 	for rows.Next() {
// 		var obj Gestor
// 		rows.Scan(&obj.ID, &obj.Nome, &obj.NomeGestor, &obj.Telefone, &obj.Login, &obj.Senha, &obj.DataCadastro, &obj.Auditoria, &obj.PrimeiroAcesso, &obj.Ativo, &obj.DataAtivo, &obj.Validado, &obj.DataValidade, &obj.Empresa.ID)
// 		ps = append(ps, Gestor{ID: obj.ID, Nome: obj.Nome, NomeGestor: obj.NomeGestor, Telefone: obj.Telefone, Login: obj.Login, Senha: obj.Senha, DataCadastro: obj.DataCadastro, Auditoria: obj.Auditoria, PrimeiroAcesso: obj.PrimeiroAcesso, Ativo: obj.Ativo, DataAtivo: obj.DataAtivo, Validado: obj.Validado, DataValidade: obj.DataValidade, Empresa: obj.Empresa})
// 	}

// 	defer rows.Close()
// 	return ps
// }

//DeleteGestor deleteFornecedor
func DeleteGestor(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var obj Gestor
	var err error
	vars := mux.Vars(req)

	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Erro ao inserir o identificador")
	}

	deletarGestor(obj.ID)

	//json.NewEncoder(w).Encode(Gestors)
}

//DeletarGestor deletar usuaria
func deletarGestor(id int) {

	stmt2, err := persistencia.DB.Prepare("delete from Gestor where id = $1")
	stmt2.Exec(id)

	if err != nil {
		panic(err)
	}
	stmt2.Close()
}

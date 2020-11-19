package modelos

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/andreggpereira/cerh--app--backend/controler"
	"github.com/andreggpereira/cerh--app--backend/persistencia"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

//Funcionario capitalize to export from package
type Funcionario struct {
	ID             int     `json:"id,omitempty"`
	Nome           string  `json:"nome,omitempty"`
	Telefone       string  `json:"telefone,omitempty"`
	Login          string  `json:"login,omitempty"`
	Senha          []byte  `json:"senha,omitempty"`
	DataCadastro   string  `json:"datacadastro,omitempty"`
	DataAdmissao   string  `json:"dataadmissao,omitempty"`
	DataNascimento string  `json:"datanascimento,omitempty"`
	Auditoria      string  `json:"auditoria,omitempty"`
	Codigo         string  `json:"codigo,omitempty"`
	PrimeiroAcesso bool    `json:"primeiroacesso,omitempty"`
	Ativo          bool    `json:"ativo,omitempty"`
	DataAtivo      string  `json:"dataativo,omitempty"`
	Validado       bool    `json:"validado,omitempty"`
	DataValidado   string  `json:"datavalidado,omitempty"`
	Sexo           Sexo    `json:"sexo,omitempty"`
	Cargo          Cargo   `json:"cargo,omitempty"`
	Setor          Setor   `json:"setor,omitempty"`
	Gestor         Gestor  `json:"gestor,omitempty"`
	Empresa        Empresa `json:"empresa,omitempty"`
}

var funcionarios []Funcionario

//FuncionarioList struct consulta
type FuncionarioList struct {
	Funcionario []Funcionario     `json:"funcionario,omitempty"`
	Contador    int               `json:"contador,omitempty"`
	Message     controler.Message `json:"message,omitempty"`
}

//JwtToken struct
type JwtToken struct {
	Token  string `json:"token"`
	Header string `json:"header"`
}

//Exception struct
type Exception struct {
	Message string `json:"message"`
}

func insertORUpFuncionario(obj Funcionario) int {

	var id int
	//Se exister ID item deverá ser atualizado
	if obj.ID != 0 {
		id = obj.ID
		stmt, err := persistencia.DB.Prepare("update funcionario set nome = $1, telefone = $2, login = $3, datacadastro = $4, dataadmissao = $5, datanascimento = $6, auditoria = $7, codigo = $8, primeiroacesso = $9, ativo = $10, dataativo = $11, validado = $12, datavalidado = $13, id_gestor = $14, id_sexo = $15, id_cargo = $16, id_setor = $17, id_empresa = $18 where id =$19")
		stmt.Exec(obj.Nome, obj.Telefone, obj.Login, obj.DataCadastro, obj.DataAdmissao, obj.DataNascimento, obj.Auditoria, obj.Codigo, obj.PrimeiroAcesso, obj.Ativo, obj.DataAtivo, obj.Validado, obj.DataValidado, obj.Sexo.ID, obj.Cargo.ID, obj.Setor.ID, obj.Gestor.ID, obj.Empresa.ID, obj.ID)
		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	} else {
		stmt, err := persistencia.DB.Prepare("insert into funcionario(nome, telefone, login, datacadastro, dataadmissao, datanascimento, auditoria, codigo, primeiroacesso, ativo, dataativo, validado, datavalidado, id_sexo, id_cargo, id_setor, id_gestor, id_empresa) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)  RETURNING id")
		err = stmt.QueryRow(obj.Nome, obj.Telefone, obj.Login, obj.DataCadastro, obj.DataAdmissao, obj.DataNascimento, obj.Auditoria, obj.Codigo, obj.PrimeiroAcesso, obj.Ativo, obj.DataAtivo, obj.Validado, obj.DataValidado, obj.Sexo.ID, obj.Cargo.ID, obj.Setor.ID, obj.Gestor.ID, obj.Empresa.ID).Scan(&id)
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

// func insertORUpFuncionario(obj Funcionario) int {

// 	var id int
// 	//Se exister ID item deverá ser atualizado
// 	if obj.ID != 0 {
// 		id = obj.ID
// 		stmt, err := persistencia.DB.Prepare("update funcionario set nome = $1, telefone = $2, login = $3, senha = $4, datacadastro = $5, dataadmissao = $6, datanascimento = $7, auditoria = $8, codigo = $9, primeiroacesso = $10, ativo = $11, dataativo = $12, validado = $13, datavalidado = $14, id_gestor = $15, id_sexo = $16, id_cargo = $17, id_setor = $18, id_empresa = $19 where id =$20")
// 		stmt.Exec(obj.Nome, obj.Telefone, obj.Login, obj.Senha, obj.DataCadastro, obj.DataAdmissao, obj.DataNascimento, obj.Auditoria, obj.Codigo, obj.PrimeiroAcesso, obj.Ativo, obj.DataAtivo, obj.Validado, obj.Datavalidado, obj.Gestor.ID, obj.Sexo.ID, obj.Cargo.ID, obj.Setor.ID, obj.Empresa.ID, obj.ID)
// 		if err != nil {
// 			log.Fatal("Não foi possível atualizar o item", err)
// 			fmt.Println(err)
// 		} else {
// 			fmt.Println("Operação concluida com sucesso!!")
// 		}
// 		defer stmt.Close()
// 	} else {

// 		stmt, err := persistencia.DB.Prepare("insert into Funcionario(nome, telefone, login, senha, datacadastro, dataadmissao, datanascimento, auditoria, codigo, primeiroacesso, ativo, dataativo, validado, datavalidado, id_gestor, id_sexo, id_cargo, id_setor, id_empresa) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)  RETURNING id")
// 		err = stmt.QueryRow(obj.Nome, obj.Telefone, obj.Login, obj.Senha, obj.DataCadastro, obj.DataAdmissao, obj.DataNascimento, obj.Auditoria, obj.Codigo, obj.PrimeiroAcesso, obj.Ativo, obj.DataAtivo, obj.Validado, obj.Datavalidado, obj.Gestor.ID, obj.Sexo.ID, obj.Cargo.ID, obj.Setor.ID, obj.Empresa.ID).Scan(&id)
// 		fmt.Println("Teste ID ", &id, id)
// 		if err != nil {
// 			log.Fatal("Cannot run insert statement", err)
// 			fmt.Println(err)
// 		} else {
// 			fmt.Println("Operação concluida com sucesso!!")
// 		}
// 		defer stmt.Close()
// 	}
// 	return id
// }

//UsuarioLogi1 login do usuario
func UsuarioLogi1(w http.ResponseWriter, req *http.Request) {

	fmt.Println(" Entou aqui no login")
	w.Header().Set("Content-Type", "application/json")

	var funcionario Funcionario
	var funcionarioDB Funcionario
	err := json.NewDecoder(req.Body).Decode(&funcionario)
	if err != nil {
		http.Error(w, "Ocorreu um erro ao realizar a operação", 403)
	}

	funcionarioDB = getFuncionarioLogin(funcionario.Login)

	if funcionarioDB.ID == 0 {
		json.NewEncoder(w).Encode(controler.Message{Message: "Login ou senha, inválidos", Status: 403})

	} else {
		//comparação do hash cadastrado no banco, com o hash da senha que foi digitada
		fmt.Println("Comprar senha hash")
		err := bcrypt.CompareHashAndPassword(funcionarioDB.Senha, []byte(funcionario.Senha))

		if err != nil {
			json.NewEncoder(w).Encode(controler.Message{Message: "Login ou senha, inválidos", Status: 403})
		} else {

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"usuario": funcionarioDB.ID,
				"empresa": funcionarioDB.Empresa.ID,
			})

			claims := token.Claims.(jwt.MapClaims)
			claims["autorized"] = true
			claims["user"] = funcionarioDB.Login
			claims["exp"] = time.Now().Add(time.Minute * 3000).Unix()

			tokenString, error := token.SignedString([]byte(controler.JwtSecretKey))

			if error != nil {
				fmt.Println(error)
			}

			json.NewEncoder(w).Encode(controler.JwtToken{Token: tokenString, Status: 202, Login: funcionarioDB.Login, Nome: funcionarioDB.Nome})
		}
	}
}

//BuscarEmail busca geral
func BuscarEmail(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Funcionario
	err := decoder.Decode(&obj)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj.Login)

	obj = getFuncionarioLogin(obj.Login)

	if obj.ID != 0 {

		json.NewEncoder(w).Encode("Email já cadastrado")
	} else {
		json.NewEncoder(w).Encode("Email válido")
	}

}

func getBuscaEmail(login string) bool {

	//conta quanto usuarios possui o banco com esse email/login
	resultado := persistencia.DB.QueryRow("select count (id) from Funcionario where login = $1", login)

	var count int
	resultado.Scan(&count)

	if count == 0 {
		return false
	}
	return true
}

func getFuncionarioLogin(login string) Funcionario {

	resultado, _ := persistencia.DB.Query("select id, nome, login, senha, id_empresa from Funcionario where login = $1", login)
	defer resultado.Close()
	var um Funcionario
	for resultado.Next() {
		resultado.Scan(&um.ID, &um.Nome, &um.Login, &um.Senha, &um.Empresa.ID)
		um = Funcionario{ID: um.ID, Nome: um.Nome, Login: um.Login, Senha: um.Senha, Empresa: um.Empresa}
	}
	fmt.Println("Daddosd usuario :", um)

	return um
}

//CreateFuncionario cadastro
func CreateFuncionario(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Funcionario
	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.DataCadastro = controler.PegarDataAtualStringNew()

	if obj.ID == 0 {
		obj.DataAdmissao = obj.DataAdmissao + " 00:00:00"
		obj.DataNascimento = obj.DataNascimento + " 00:00:00"
		obj.DataAtivo = "1970-01-01 00:00:01"
		obj.DataValidado = "1970-01-01 00:00:01"
		obj.PrimeiroAcesso = false
		obj.Codigo = "123123123"
		// gerar codigo de acesso
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)

	insertORUpFuncionario(obj)

	json.NewEncoder(w).Encode("Usuario cadastrado com sucesso")
}

//EnviarEmailToken busca geral
func EnviarEmailToken(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Entrou aqui kkkkkk  ")
	w.Header().Set("Content-Type", "application/json")
	var obj Funcionario
	err := json.NewDecoder(req.Body).Decode(&obj)

	fmt.Println("OBJ", obj)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("sdasdasdas ", obj.Login)

	obj = getFuncionarioLogin(obj.Login)

	fmt.Println("Dados asdsadsa ", obj)
	if obj.ID != 0 {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"usuario": obj.ID,
			"empresa": obj.Empresa.ID,
		})

		claims := token.Claims.(jwt.MapClaims)

		claims["autorized"] = true
		claims["user"] = obj.Login
		claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

		tokenString, error := token.SignedString([]byte("sisalicerce"))
		if error != nil {
			fmt.Println(error)
		}
		controler.EnviarEmail(obj.Login, tokenString)
		json.NewEncoder(w).Encode(JwtToken{Token: tokenString, Header: "authorization"})
		//json.NewEncoder(w).Encode(JwtToken{Token: tokenString}, controle.Message{Header: "authorization", StatusCode: 200})

	} else {
		json.NewEncoder(w).Encode("Usuario não encontrado")
	}

}

//EnviarEmailTokenAcesso busca geral
func EnviarEmailTokenAcesso(idusuario, impresa int, login string) {
	if idusuario != 0 {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"usuario": idusuario,
			"empresa": impresa,
		})

		claims := token.Claims.(jwt.MapClaims)

		claims["autorized"] = true
		claims["user"] = login
		claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

		tokenString, error := token.SignedString([]byte("sisalicerce"))
		if error != nil {
			fmt.Println(error)
		}

		controler.EnviarEmail(login, tokenString)
		fmt.Println("Email enviado com sucesso")
		//json.NewEncoder(w).Encode(JwtToken{Token: tokenString, Header: "authorization"})
		//json.NewEncoder(w).Encode(JwtToken{Token: tokenString}, controle.Message{Header: "authorization", StatusCode: 200})

	} else {
		fmt.Println("Não foi possivel gerar o token para acesso")
	}

}

//ValidarEmailToken token
func ValidarEmailToken(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Entrou aqui kkkkkk  ")
	w.Header().Set("Content-Type", "application/json")
	pegartoken := req.URL.Query()
	token := pegartoken["token"]
	var tokenString = token[0]

	authorizationHeader := "Bearer " + tokenString
	fmt.Println("authorizationHeader ", authorizationHeader)

	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		fmt.Println("bearerToken llll ", bearerToken)
		if len(bearerToken) == 2 {
			fmt.Println("entrou aqui ")
			token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Token invalido")
				}
				return controler.JwtSecretKey, nil
			})

			if error != nil {
				json.NewEncoder(w).Encode("Token invalido")
				return
			}
			if token.Valid {

				empresa := fmt.Sprintf("%v", token.Claims.(jwt.MapClaims)["empresa"])
				usuario := fmt.Sprintf("%v", token.Claims.(jwt.MapClaims)["usuario"])

				req.Header.Set("empresa", empresa)
				req.Header.Set("usuario", usuario)
				json.NewEncoder(w).Encode(controler.Message{Message: "authorization", Status: 202})
			} else {
				json.NewEncoder(w).Encode(Exception{Message: "Token invalido"})
			}
		}
	} else {
		fmt.Println("Voce nao tem permissão")
		json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString, Header: "authorization"})

}

//ValidateToken validado
func ValidateToken(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(controler.JwtTokenValid{Valid: true})
}

//GetFuncionarioAll pegar todos os itens
func GetFuncionarioAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(req)
	//var page int
	var obj Funcionario
	var objList FuncionarioList
	var err error
	obj.Empresa.ID, _ = controler.PegarIDEmpresaIDGestor(req)

	// /page = controler.StringForInt(vars["page"])

	objList.Funcionario, err = getFuncionarioALL(obj.Empresa.ID)
	objList.Contador = len(objList.Funcionario)

	if err != nil {
		objList.Message.Message = "Nao possivel realizar a consulta"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 200

	}

	fmt.Println(" Dados da consulta ", objList)

	json.NewEncoder(w).Encode(&objList)
}

func getFuncionarioALL(idempresa int) ([]Funcionario, error) {
	rows, err := persistencia.DB.Query("SELECT id, nome, telefone, login, datacadastro, dataadmissao, datanascimento, auditoria, codigo, primeiroacesso, ativo, dataativo, validado, id_gestor, id_sexo, id_cargo, id_setor, id_empresa FROM funcionario WHERE id_empresa =$1 ORDER BY nome ASC", idempresa)

	if err != nil {
		fmt.Println(err)
	}

	var ps []Funcionario
	for rows.Next() {
		var obj Funcionario

		rows.Scan(&obj.ID, &obj.Nome, &obj.Telefone, &obj.Login, &obj.DataCadastro, &obj.DataAdmissao, &obj.DataNascimento, &obj.Auditoria, &obj.Codigo, &obj.PrimeiroAcesso, &obj.Ativo, &obj.DataAtivo, &obj.Validado, &obj.Gestor.ID, &obj.Sexo.ID, &obj.Cargo.ID, &obj.Setor.ID, &obj.Empresa.ID)
		ps = append(ps, Funcionario{ID: obj.ID, Nome: obj.Nome, Telefone: obj.Telefone, Login: obj.Login, DataCadastro: obj.DataCadastro, DataAdmissao: obj.DataAdmissao, DataNascimento: obj.DataNascimento, Auditoria: obj.Auditoria, Codigo: obj.Codigo, PrimeiroAcesso: obj.PrimeiroAcesso, Ativo: obj.Ativo, DataAtivo: obj.DataAtivo, Validado: obj.Validado, Gestor: obj.Gestor, Sexo: obj.Sexo, Cargo: obj.Cargo, Setor: obj.Setor, Empresa: obj.Empresa})
	}

	defer rows.Close()
	return ps, err
}

//GetFuncionarioPGAll pegar todos os itens
func GetFuncionarioPGAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var page int
	var obj Funcionario
	var objList FuncionarioList
	var err error

	obj.Empresa.ID, _ = controler.PegarIDEmpresaIDGestor(req)
	page = controler.StringForInt(vars["page"])

	objList.Funcionario, err = getFuncionaioPGAll(obj.Empresa.ID, page)
	objList.Contador = getCountFuncinarioAll(obj.Empresa.ID)

	if err != nil {
		objList.Message.Message = "Nao possivel realizar a consulta"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 200

	}

	json.NewEncoder(w).Encode(&objList)
}

func getFuncionaioPGAll(idempresa, page int) ([]Funcionario, error) {
	rows, err := persistencia.DB.Query("SELECT id, nome, telefone, login, datacadastro, dataadmissao, datanascimento, auditoria, codigo, primeiroacesso, ativo, dataativo, validado, id_sexo, id_cargo, id_setor, id_gestor FROM funcionario WHERE id_empresa =$1 ORDER BY nome DESC LIMIT $2 OFFSET $3 ", idempresa, 10, page)

	if err != nil {
		fmt.Println(err)
	}

	var ps []Funcionario
	for rows.Next() {
		var obj Funcionario

		rows.Scan(&obj.ID, &obj.Nome, &obj.Telefone, &obj.Login, &obj.DataCadastro, &obj.DataAdmissao, &obj.DataNascimento, &obj.Auditoria, &obj.Codigo, &obj.PrimeiroAcesso, &obj.Ativo, &obj.DataAtivo, &obj.Validado, &obj.Sexo.ID, &obj.Cargo.ID, &obj.Setor.ID, &obj.Gestor.ID)
		ps = append(ps, Funcionario{ID: obj.ID, Nome: obj.Nome, Telefone: obj.Telefone, Login: obj.Login, DataCadastro: obj.DataCadastro, DataAdmissao: obj.DataAdmissao, DataNascimento: obj.DataNascimento, Auditoria: obj.Auditoria, Codigo: obj.Codigo, PrimeiroAcesso: obj.PrimeiroAcesso, Ativo: obj.Ativo, DataAtivo: obj.DataAtivo, Validado: obj.Validado, Sexo: obj.Sexo, Cargo: obj.Cargo, Setor: obj.Setor, Gestor: obj.Gestor})
	}

	defer rows.Close()
	return ps, err
}

func getCountFuncinarioAll(idempresa int) int {

	rows, err := persistencia.DB.Query("SELECT COUNT(id) FROM funcionario where id_empresa = $1 ", idempresa)

	if err != nil {
		fmt.Println(err)
	}
	var ps []Funcionario
	for rows.Next() {
		var um Funcionario
		rows.Scan(&um.ID)
		ps = append(ps, Funcionario{ID: um.ID})
	}

	defer rows.Close()

	var contador int
	for _, numero := range ps {
		contador = numero.ID
	}
	return contador
}

// func getFuncionarioALL(idempresa int) ([]Funcionario, error) {
// 	rows, err := persistencia.DB.Query("SELECT id, nome, telefone, login, datacadastro, dataadmissao, datanascimento, auditoria, codigo, primeiroacesso, ativo, dataativo, validado id_gestor, id_sexo, id_cargo, id_setor, id_empresa FROM funcionario WHERE id_empresa =$1 ORDER BY nome ASC", idempresa)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var ps []Funcionario
// 	for rows.Next() {
// 		var obj Funcionario

// 		rows.Scan(&obj.ID, &obj.Nome, &obj.Telefone, &obj.Login, &obj.DataCadastro, &obj.DataAdmissao, &obj.DataNascimento, &obj.Auditoria, &obj.Codigo, &obj.PrimeiroAcesso, &obj.Ativo, &obj.DataAtivo, &obj.Validado, &obj.Gestor.ID, &obj.Sexo.ID, &obj.Cargo.ID, &obj.Setor.ID, &obj.Empresa.ID)
// 		ps = append(ps, Funcionario{ID: obj.ID, Nome: obj.Nome, Telefone: obj.Telefone, Login: obj.Login, DataCadastro: obj.DataCadastro, DataAdmissao: obj.DataAdmissao, DataNascimento: obj.DataNascimento, Auditoria: obj.Auditoria, Codigo: obj.Codigo, PrimeiroAcesso: obj.PrimeiroAcesso, Ativo: obj.Ativo, DataAtivo: obj.DataAtivo, Validado: obj.Validado, Gestor: obj.Gestor, Sexo: obj.Sexo, Cargo: obj.Cargo, Setor: obj.Setor, Empresa: obj.Empresa})
// 	}

// 	defer rows.Close()
// 	return ps, err
// }

// func getFuncionaioPGAll(idempresa int, page int) ([]Funcionario, error) {
// 	rows, err := persistencia.DB.Query("SELECT id, nome, datacadastro, email, telefone, telefonesecu, cpf, cnpj, pessoafisica, obspessoa, id_usuario, cep, endereco, uf, cidade, bairro, numero, complemento FROM cliente WHERE id_empresa =$1 ORDER BY nome DESC LIMIT $2 OFFSET $3 ", idempresa, 10, page)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	var ps []Funcionario
// 	for rows.Next() {
// 		var obj Funcionario

// 		rows.Scan(&obj.ID, &obj.Nome, &obj.DataCadastro)
// 		ps = append(ps, Funcionario{ID: obj.ID, Nome: obj.Nome, DataCadastro: obj.DataCadastro})
// 	}

// 	defer rows.Close()
// 	return err, ps
// }

//DeleteFuncionario deleteFornecedor
func DeleteFuncionario(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var obj Funcionario
	var err error
	vars := mux.Vars(req)

	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Erro ao inserir o identificador")
	}

	DeletarFuncionario(obj.ID)

	json.NewEncoder(w).Encode(funcionarios)
}

//DeletarFuncionario deletar usuaria
func DeletarFuncionario(id int) {

	stmt2, err := persistencia.DB.Prepare("delete from Funcionario where id = $1")
	stmt2.Exec(id)

	if err != nil {
		panic(err)
	}
	stmt2.Close()
}

/*

type session struct {
	un           string
	lastActivity time.Time
}
	func login(w http.ResponseWriter, req *http.Request) {

		if req.Method == http.MethodPost {
			var u modelos.Usuario
			un := req.FormValue("login")
			p := req.FormValue("senha")

			//verficica se a algum usuario com o login inserido
			logi, arr := modelos.GetUsuarioLogin(un)
			if logi == false {
				http.Error(w, "Senha e Login Inválidos", http.StatusForbidden)
				return
			}

			for i, dado := range arr {
				u = dado
				fmt.Println("Informação", dado)
				fmt.Println("Indice", i)
				u, ok := dbUsers[un]

				fmt.Println("Dados =", u)
				fmt.Println("Dados OKc=", ok)
			}
			//comparação do hash cadastrado no banco, com o hash da senha que foi digitada
			err := bcrypt.CompareHashAndPassword(u.Senha, []byte(p))
			if err != nil {
				http.Error(w, "Login ou senha, inválidos", http.StatusForbidden)
				return
			}
			// criar a sessao
			sID, _ := uuid.NewV4()
			c := &http.Cookie{
				Name:  "session",
				Value: sID.String(),
			}
			c.MaxAge = sessionLength
			http.SetCookie(w, c)
			dbSessions[c.Value] = session{un, time.Now()}

			u.Senha = nil
			http.Redirect(w, req, "/usuario", http.StatusMovedPermanently)

			//imprimir dados da session
			showSessions()
			return
		}
		showSessions()
		http.Redirect(w, req, "/", http.StatusMovedPermanently)
	}

func UsuarioLogin(w http.ResponseWriter, req *http.Request) {

	var usuario Usuario
	var usuarioDB Usuario

	decoder := json.NewDecoder(req.Body)

	erro := decoder.Decode(&usuario)
	if erro != nil {
		panic(erro)
	}

	fmt.Println("Usuario digitado", usuario)
	fmt.Println("Usuario digitado", usuario.Login)
	fmt.Println("Senha digitado", string(usuario.Senha))

	//verficica se a algum usuario com o login inserido
	logi, arr := GetUsuarioLogin(usuario.Login)
	if logi == false {
		http.Error(w, "Senha e Login Inválidos", http.StatusForbidden)
		return
	}

	for i, dado := range arr {
		usuarioDB = dado
		usuarioDB, ok := dbUsers[usuarioDB.Login]
	}
	//comparação do hash cadastrado no banco, com o hash da senha que foi digitada
	err := bcrypt.CompareHashAndPassword(usuarioDB.Senha, []byte(usuario.Senha))
	//err := bcrypt.CompareHashAndPassword(u.Senha, []byte(p))
	if err != nil {
		http.Error(w, "Login ou senha, inválidos", http.StatusForbidden)
		return
	}
	// criar a sessao
	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	dbSessions[c.Value] = session{usuarioDB.Login, time.Now()}

	ShowSessions()
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

//CadastroUsuario -Create
func CadastroUsuario(user Usuario) {

	stmt, _ := persistencia.DB.Prepare("insert into usuario(nome,email,senha,id_permissao) values(?, ?, ?, ?)")
	res, err := stmt.Exec(user.Nome, user.Login, user.Senha, 1)

	fmt.Println(err)
	log.Panic(err)

	id, err := res.LastInsertId()
	fmt.Println("Ultimo id inserido", id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}

func CadUsuario(Usuario Usuario) {

	if Usuario.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update Usuario set nome = $1,login = $2, senha = $3, id_permissao = $4 where id =$5")
		stmt.Exec(Usuario.Nome, Usuario.Login, Usuario.Senha, Usuario.Permissao.ID, Usuario.ID)

		if err != nil {
			log.Fatal("Cannot run insert statement", err)
			panic(err)
			fmt.Println("Operação não concluida!!")
			//	fmt.Println("Res", res)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}

	} else {

		stmt, err := persistencia.DB.Prepare("insert into Usuario(nome,login,senha,id_permissao) values($1,$2,$3,$4)")
		res, err := stmt.Exec(Usuario.Nome, Usuario.Login, Usuario.Senha, Usuario.Permissao.ID)

		if err != nil {
			log.Fatal("Cannot run insert statement", err)
			panic(err)
			fmt.Println("Operação não concluida!!")
			fmt.Println("Res", res)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
	}
}

//Lista de todos os estados
func GetUsuarioAll() []Usuario {

	rows, err := persistencia.DB.Query("SELECT us.id, us.nome, us.login,us.senha, us.id_permissao, per.nome FROM usuario AS us INNER JOIN permissao AS per ON us.id_permissao = per.id")

	fmt.Println("Rows : ", rows)
	fmt.Println("Erro : ", err)

	var ps []Usuario
	for rows.Next() {
		var um Usuario
		rows.Scan(&um.ID, &um.Nome, &um.Login, &um.Senha, &um.Permissao.ID, &um.Permissao.Nome)
		ps = append(ps, Usuario{ID: um.ID, Nome: um.Nome, Login: um.Login, Senha: um.Senha, Permissao: um.Permissao})
	}
	return ps
}
func DeletarUsuario(id int) {

	stmt2, err := persistencia.DB.Prepare("delete from usuario where id = $1")
	stmt2.Exec(id)

	if err != nil {
		panic(err)
	}
	stmt2.Close()
}
func GetUsuarioLogin(login string) (bool, []Usuario) {

	resultado, _ := persistencia.DB.Query("select id,nome,login,senha from usuario where login = $1", login)
	fmt.Println("Buscou por login")
	var ps []Usuario
	for resultado.Next() {
		var um Usuario
		resultado.Scan(&um.ID, &um.Nome, &um.Login, &um.Senha)
		ps = append(ps, Usuario{ID: um.ID, Nome: um.Nome, Login: um.Login, Senha: um.Senha})
	}

	if len(ps) > 0 {
		return true, ps
	} else {
		return false, ps
	}
}
func GetUsuarioID(id int) (bool, []Usuario) {

	resultado, _ := persistencia.DB.Query("SELECT us.id, us.nome, us.login,us.senha, us.id_permissao, per.nome FROM usuario AS us INNER JOIN permissao AS per ON us.id_permissao = per.id AND us.id = $1", id)
	var ps []Usuario
	for resultado.Next() {
		var um Usuario
		resultado.Scan(&um.ID, &um.Nome, &um.Login, &um.Senha, &um.Permissao.ID, &um.Permissao.Nome)
		ps = append(ps, Usuario{ID: um.ID, Nome: um.Nome, Login: um.Login, Senha: um.Senha, Permissao: um.Permissao})
	}
	if len(ps) > 0 {
		return true, ps
	} else {
		return false, ps
	}
}
func pegarEstado1ID(idEstado int64) *Estado {

	db, err := sql.Open("postgres", "postgres://postgres:110407@localhost/racp?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// consulta usuario
	rows, _ := db.Query("select id,nome,datacadastro from estado where id > ?", idEstado)
	defer rows.Close()

	// for para percorrer as linhas

	for rows.Next() {
		var u Estado
		rows.Scan(&u.ID, &u.Nome)
		fmt.Println(u)
	}
	return nil
}
func CreateUsuario(w http.ResponseWriter, req *http.Request) {

	// verifica se o usuario esta logado
	//	if (alreadyLoggedIn(w, req)) == false {
	//		fmt.Println("Deu false")
	//		http.Redirect(w, req, "/", http.StatusSeeOther)
	//		return
	//	}

	//teste(w, req)

	var Usuario Usuario
	if req.Method == http.MethodPost {

		id := req.FormValue("id")
		un := req.FormValue("nome")
		l := req.FormValue("login")
		p := req.FormValue("senha")

		//converte campo do formulario em int
		Usuario.ID, _ = strconv.Atoi(id)

		//verifica se campor login esta vazio, senão busta o login
		if l != "" {
			ok, arr := GetUsuarioLogin(l)

			// se o resultado da busca for true
			if ok == true {
				for _, dados := range arr {

					if dados.ID == Usuario.ID {
						ok = false
					}
				}
			}

			if ok == true {
				http.Error(w, "Username and/or password do not match", http.StatusForbidden)
				return

			} else {

				if s, err := strconv.ParseInt(req.FormValue("permissao"), 10, 64); err == nil {
					Usuario.Permissao.ID = int(s)
				}

				//criando a session
				sID, _ := uuid.NewV4()
				c := &http.Cookie{
					Name:  "session",
					Value: sID.String(),
				}
				c.MaxAge = sessionLength
				http.SetCookie(w, c)
				dbSessions[c.Value] = session{un, time.Now()}

				//gerar hash da senha
				//sum := sha256.Sum256([]byte())
				bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
				if err != nil {
					http.Error(w, "Internal server error", http.StatusInternalServerError)
					return
				}
				//popula a struct
				Usuario.ID, _ = strconv.Atoi(id)
				Usuario.Nome = un
				Usuario.Login = l
				Usuario.Senha = bs
				dbUsers[un] = Usuario

				CadUsuario(Usuario)
			}
		}
	}
}
func delusuario(w http.ResponseWriter, req *http.Request) {

	if !AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if s, err := strconv.ParseInt(req.FormValue("id"), 10, 64); err == nil {
		DeletarUsuario(int(s))
	} else {
	}
	http.Redirect(w, req, "/usuario", http.StatusMovedPermanently)
}
func altusuario(w http.ResponseWriter, req *http.Request) {

	fmt.Println("AltAusuairo")
	if (AlreadyLoggedIn(w, req)) == false {
		fmt.Println("Deu false")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var Usuario Usuario
	// converte string em int
	idUs, _ := strconv.Atoi(req.FormValue("id"))
	//busca os usuario por id
	ok, usuarios := GetUsuarioID(idUs)

	//verifica se a usuarios na
	if ok == false {
		http.Error(w, "Operação não Realizada, erro ao alterar o usuario", http.StatusForbidden)
		return
	}
	//popula struct
	for _, dado := range usuarios {
		Usuario = dado
	}

	Usuario.Senha = nil

	//	p1 := TodoPageData{Usuarios: modelos.GetUsuarioAll(), Permissao: modelos.GetPermissaoALL(), Usuario: Usua}
	controler.TPL.ExecuteTemplate(w, "altusuario.html", nil)
	//tpl.ExecuteTemplate(w, "usuario.html", Usuario)
}

//GetUser pegar usuario
func GetUser(w http.ResponseWriter, req *http.Request) Usuario {

	//tenta pegar o Cookie
	c, err := req.Cookie("session")
	// se não achou, cria novo
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// se usuario existe, retorna o usuario
	var u Usuario
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

//Metado  serve para testar se o usuario esta logado
func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {

	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	s, ok := dbSessions[c.Value]

	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}

//CleanSessions limpar Session
func CleanSessions() {
	fmt.Println("BEFORE CLEAN")
	ShowSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	ShowSessions()
}

//ShowSessions imprimir dados da session
func ShowSessions() {
	fmt.Println("********")
	fmt.Println("Dados da Session")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
*/

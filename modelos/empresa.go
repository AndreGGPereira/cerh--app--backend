package modelos

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/andreggpereira/cerh--app--backend/controler"
	"github.com/andreggpereira/cerh--app--backend/persistencia"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

//Empresa  CRUD
type Empresa struct {
	ID           int    `json:"id,omitempty"`
	RazaoSocial  string `json:"razaosocial,omitempty"`
	DataCadastro string `json:"datacadastro,omitempty"`
	Email        string `json:"email,omitempty"`
	EmailSecu    string `json:"emailsecu,omitempty"`
	Telefone     string `json:"telefone,omitempty"`
	TelefoneSecu string `json:"telefonesecu,omitempty"`
	CPF          string `json:"cpf,omitempty"`
	CNPJ         string `json:"cnpj,omitempty"`
	ObsPessoa    string `json:"obspessoa,omitempty"`
	PessoaFisica bool   `json:"pessoafisica,omitempty"`
	Acesso       bool   `json:"acesso,omitempty"`
}

var empresas []Empresa

//CreateEmpresa cadastrar e atualizar
func CreateEmpresa(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	payload := make(map[string]interface{})
	err := decoder.Decode(&payload)
	fmt.Println(err)

	fmt.Println(" Entrou aqui --------")
	var obj Empresa
	var gestor Gestor
	var enderecoEmpresa EnderecoEmpresa

	for i, v := range payload {
		fmt.Println(i, "=", v)

		switch i {
		case "id":
			obj.ID, _ = v.(int)
			continue
		case "razaosocial":
			obj.RazaoSocial, _ = v.(string)
			continue
		case "datacadastro":
			obj.DataCadastro, _ = v.(string)
			continue
		case "email":
			obj.Email, _ = v.(string)
			continue
		case "emailsecu":
			obj.EmailSecu, _ = v.(string)
			continue
		case "telefone":
			obj.Telefone, _ = v.(string)
			continue
		case "cpf":
			obj.CPF, _ = v.(string)
			continue
		case "cnpj":
			obj.CNPJ, _ = v.(string)
			continue
		case "obspessoa":
			obj.ObsPessoa, _ = v.(string)
			continue
		case "pessoadisica":
			obj.PessoaFisica, _ = v.(bool)
			continue
		case "acesso":
			obj.Acesso, _ = v.(bool)
			continue
		case "cep":
			enderecoEmpresa.Cep, _ = v.(string)
			continue
		case "endereco":
			enderecoEmpresa.Endereco, _ = v.(string)
			continue
		case "bairro":
			enderecoEmpresa.Bairro, _ = v.(string)
			continue
		case "cidade":
			enderecoEmpresa.Cidade, _ = v.(string)
			continue
		case "complemento":
			enderecoEmpresa.Complemento, _ = v.(string)
			continue
		case "numero":
			fmt.Println("numero :", v.(string))
			enderecoEmpresa.Numero, _ = strconv.Atoi(v.(string))
			continue
		case "uf":
			enderecoEmpresa.Uf, _ = v.(string)
			continue
		case "nome":
			gestor.Nome, _ = v.(string)
			continue
		case "login":
			gestor.Login, _ = v.(string)
			continue
		case "senha":
			fmt.Println("Senha :", v.(string))
			gestor.Senha = []byte(v.(string))
			fmt.Println("Senha :", gestor.Senha)
			continue
		}
	}

	obj.Email = "andreggp@gmail.com"

	obj.DataCadastro = controler.PegarDataAtualString()
	fmt.Println(" Dados do obj ", obj)
	fmt.Println(" Dados do usuario ", gestor)
	fmt.Println(" Dados do enderecoEmpresa ", enderecoEmpresa)

	obj.ID = insertORUpEmpresa(obj)

	enderecoEmpresa.Empresa.ID = obj.ID
	gestor.Empresa.ID = obj.ID

	fmt.Println("Senha dadis di gestor :", gestor)
	fmt.Println("Senha  obj.ID :", obj.ID)
	fmt.Println("Senha  obj :", obj)
	if obj.ID != 0 {
		// 	insertORUpEnderecoEmpresa(enderecoEmpresa)

		fmt.Println("Usuario Senha === ", []byte(gestor.Senha))
		hashSenha, err := bcrypt.GenerateFromPassword([]byte(gestor.Senha), bcrypt.MinCost)
		if err != nil {
			fmt.Println("Erro ao gerar o Hash")
		}

		gestor.Senha = hashSenha
		gestor.Login = obj.Email
		gestor.PrimeiroAcesso = true
		gestor.Telefone = obj.Telefone
		gestor.Empresa.ID = obj.ID
		gestor.DataCadastro = controler.PegarDataAtualStringNew()
		gestor.DataAtivo = "1970-01-01 00:00:01"
		gestor.DataValidade = "1970-01-01 00:00:01"

		fmt.Println("Senha  obj :", gestor)

		//gestor.Senha = hashSenha
		gestor.Empresa.ID = insertORUpGestor(gestor)
		//primeiro acesso empresa

		// 	// primerioAcessoDadosPermissaoTipo(gestor.ID, obj.ID)
		// 	// //mandar email com token
		//EnviarEmailTokenAcesso(gestor.ID, obj.ID, gestor.Login)
	} else {
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}
}

func insertORUpEmpresa(obj Empresa) int {
	var id int
	//Se exister ID item deverá ser atualizado
	if obj.ID != 0 {
		stmt, err := persistencia.DB.Prepare("update Empresa set RazaoSocial = $1, datacadastro = $2, email = $3, emailsecu = $4, telefone = $5, telefonesecu = $6, cpf = $7, cnpj = $8, pessoafisica = $9, obspessoa = $10, acesso = $11  where id =$12 ")
		stmt.Exec(obj.RazaoSocial, obj.DataCadastro, obj.Email, obj.EmailSecu, obj.Telefone, obj.CPF, obj.CNPJ, obj.PessoaFisica, obj.ObsPessoa, obj.Acesso, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {
		obj.PessoaFisica = true
		obj.Acesso = true
		stmt, err := persistencia.DB.Prepare("insert into empresa(razaosocial, datacadastro, email, emailsecu, telefone, telefonesecu, cpf, cnpj, pessoafisica, obspessoa, acesso ) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id")

		err = stmt.QueryRow(obj.RazaoSocial, obj.DataCadastro, obj.Email, obj.EmailSecu, obj.Telefone, obj.TelefoneSecu, obj.CPF, obj.CNPJ, obj.PessoaFisica, obj.ObsPessoa, obj.Acesso).Scan(&id)
		fmt.Println("Teste ID ", &id, id)
		if err != nil {
			log.Fatal("Cannot run insert statement", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	}
	return id
}

//DeleteEmpresa deletar empresa
func DeleteEmpresa(w http.ResponseWriter, req *http.Request) {

	id, erro := strconv.Atoi(req.FormValue("id"))

	if id == 0 || erro != nil {
		fmt.Println("Não foi possível identificar a permissão")
	}
	deletarEmpresa(id)

}
func deletarEmpresa(id int) {

	stmt, err := persistencia.DB.Prepare("delete from empresa where id = $1")
	stmt.Exec(id)

	if err != nil {
		panic(err)
	}
	stmt.Close()
}

//GetEmpresa empresa
func GetEmpresa(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range empresas {
		if item.ID == controler.StringForInt(params["id"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Empresa{})

}
func getEmpresaID(ID int) Empresa {

	rows, _ := persistencia.DB.Query("select * from Empresa where id = $1", ID)

	var ps []Empresa
	for rows.Next() {
		var ot Empresa
		rows.Scan(&ot.ID, &ot.RazaoSocial, &ot.DataCadastro, &ot.Email, &ot.EmailSecu, &ot.Telefone, &ot.CPF, &ot.CNPJ, &ot.PessoaFisica, &ot.ObsPessoa, &ot.Acesso)
		ps = append(ps, Empresa{ID: ot.ID, RazaoSocial: ot.RazaoSocial, DataCadastro: ot.DataCadastro, Email: ot.Email, EmailSecu: ot.EmailSecu, Telefone: ot.Telefone, CPF: ot.CPF, CNPJ: ot.CNPJ, PessoaFisica: ot.PessoaFisica, ObsPessoa: ot.ObsPessoa, Acesso: ot.Acesso})
	}
	defer rows.Close()
	var obj Empresa
	for _, numero := range ps {
		obj = numero
	}
	return obj
}

//GetEmpresaAll empresa todos
func GetEmpresaAll(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	empresas, erro := getEmpresaAll()

	if erro != nil {
		panic(erro)
	}

	json.NewEncoder(w).Encode(&empresas)

}
func getEmpresaAll() ([]Empresa, error) {
	rows, err := persistencia.DB.Query("SELECT id, nome, datacadastro, email, emailsecu, telefone, telefonesecu, cpf, cnpj, pessoafisica, obspessoa FROM empresa")

	if err != nil {
		panic(err)
	}

	var ps []Empresa
	for rows.Next() {
		var ot Empresa
		rows.Scan(&ot.ID, &ot.RazaoSocial, &ot.DataCadastro, &ot.Email, &ot.EmailSecu, &ot.Telefone, &ot.CPF, &ot.CNPJ, &ot.PessoaFisica, &ot.ObsPessoa, &ot.Acesso)
		ps = append(ps, Empresa{ID: ot.ID, RazaoSocial: ot.RazaoSocial, DataCadastro: ot.DataCadastro, Email: ot.Email, EmailSecu: ot.EmailSecu, Telefone: ot.Telefone, CPF: ot.CPF, CNPJ: ot.CNPJ, PessoaFisica: ot.PessoaFisica, ObsPessoa: ot.ObsPessoa, Acesso: ot.Acesso})
	}

	defer rows.Close()
	return ps, err
}

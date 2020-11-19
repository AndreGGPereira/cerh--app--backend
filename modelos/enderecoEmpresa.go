package modelos

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andreggpereira/sisalicerce--app--backend/controler"
	"github.com/andreggpereira/sisalicerce--app--backend/persistencia"
	"github.com/gorilla/mux"
)

//EnderecoEmpresa endereco empresa
type EnderecoEmpresa struct {
	ID          int     `json:"id,omitempty"`
	Cep         string  `json:"cep,omitempty"`
	Endereco    string  `json:"endereco,omitempty"`
	Bairro      string  `json:"bairro,omitempty"`
	Cidade      string  `json:"cidade,omitempty"`
	Complemento string  `json:"complemento,omitempty"`
	Numero      int     `json:"numero,omitempty"`
	Uf          string  `json:"uf,omitempty"`
	Ddd         string  `json:"ddd,omitempty"`
	Unidade     string  `json:"unidade,omitempty"`
	Ibge        string  `json:"ibge,omitempty"`
	Empresa     Empresa `json:"empresa,omitempty"`
}

var enderecosEmpresa []EnderecoEmpresa

//CreateEnderecoEmpresa cadastrar endereco
func CreateEnderecoEmpresa(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj EnderecoEmpresa
	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	insertORUpEnderecoEmpresa(obj)
	json.NewEncoder(w).Encode("Endereço cadastrado com sucesso")
}

func insertORUpEnderecoEmpresa(obj EnderecoEmpresa) {

	//Se exister ID item deverá ser atualizado
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update EnderecoEmpresa set cep = $1, endereco = $2, bairro = $3, complemento = $4, numero = $5 cidade = $6, uf = $7, ddd = $8, unidade = $9, ibge = $10, id_empresa = $11 where id =$12")
		stmt.Exec(obj.Cep, obj.Endereco, obj.Bairro, obj.Complemento, obj.Cidade, obj.Uf, obj.Ddd, obj.Unidade, obj.Ibge, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into EnderecoEmpresa(cep, endereco, bairro, complemento, numero, cidade, uf, id_empresa) values($1, $2, $3, $4, $5, $6, $7, $8)")
		stmt.Exec(obj.Cep, obj.Endereco, obj.Bairro, obj.Complemento, obj.Numero, obj.Cidade, obj.Uf, obj.Empresa.ID)

		if err != nil {
			log.Fatal("Cannot run insert statement", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	}
}

//DeleteEnderecoEmpresa empresa
func DeleteEnderecoEmpresa(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range enderecosEmpresa {
		if item.ID == controler.StringForInt(params["id"]) {
			enderecosEmpresa = append(enderecosEmpresa[:index], enderecosEmpresa[index+1:]...)
			deletarEnderecoEmpresa(controler.StringForInt(params["id"]))
			json.NewEncoder(w).Encode("Estado Removido com Sucesso")
			break
		}
	}
	json.NewEncoder(w).Encode(enderecosEmpresa)
}
func deletarEnderecoEmpresa(id int) {

	stmt, err := persistencia.DB.Prepare("delete from EnderecoEmpresa where id = $1")
	stmt.Exec(id)

	if err != nil {
		panic(err)
	}
	stmt.Close()
}

//GetEnderecoEmpresa empresa
func GetEnderecoEmpresa(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range enderecosEmpresa {
		if item.ID == controler.StringForInt(params["id"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&EnderecoEmpresa{})
}
func getEnderecoEmpresaID(ID int) EnderecoEmpresa {

	rows, _ := persistencia.DB.Query("select id, cep, endereco, bairro, complemento, cidade, uf, ddd, unidade, ibge, id_empresa FROM EnderecoEmpresa where id_empresa = $1", ID)

	var ps []EnderecoEmpresa
	for rows.Next() {
		var e EnderecoEmpresa
		rows.Scan(&e.ID, &e.Cep, &e.Endereco, &e.Bairro, &e.Complemento, &e.Cidade, &e.Uf, &e.Ddd, &e.Unidade, &e.Ibge, &e.Empresa.ID)
		ps = append(ps, EnderecoEmpresa{ID: e.ID, Cep: e.Cep, Endereco: e.Endereco, Bairro: e.Bairro, Complemento: e.Complemento, Cidade: e.Cidade, Uf: e.Uf, Ddd: e.Ddd, Unidade: e.Unidade, Ibge: e.Ibge, Empresa: e.Empresa})
	}
	defer rows.Close()
	var obj EnderecoEmpresa
	for _, numero := range ps {
		obj = numero
	}
	return obj
}

//GetEnderecoEmpresaAll empresa
func GetEnderecoEmpresaAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	enderecosEmpresa, erro := getEnderecoEmpresaAll()

	if erro != nil {
		panic(erro)
	}

	json.NewEncoder(w).Encode(&enderecosEmpresa)

}
func getEnderecoEmpresaAll() ([]EnderecoEmpresa, error) {

	rows, err := persistencia.DB.Query("select id, cep, endereco, bairro, complemento, cidade, uf, ddd, unidade, ibge,  id_empresa  FROM enderecoempresa")

	if err != nil {
		panic(err)
	}

	var ps []EnderecoEmpresa
	for rows.Next() {
		var e EnderecoEmpresa
		rows.Scan(&e.ID, &e.Cep, &e.Endereco, &e.Bairro, &e.Complemento, &e.Cidade, &e.Uf, &e.Ddd, &e.Unidade, &e.Ibge, &e.Empresa.ID)
		ps = append(ps, EnderecoEmpresa{ID: e.ID, Cep: e.Cep, Endereco: e.Endereco, Bairro: e.Bairro, Complemento: e.Complemento, Cidade: e.Cidade, Uf: e.Uf, Ddd: e.Ddd, Unidade: e.Unidade, Ibge: e.Ibge, Empresa: e.Empresa})

	}
	defer rows.Close()
	return ps, err
}

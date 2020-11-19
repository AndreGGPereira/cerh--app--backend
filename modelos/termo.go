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
)

//Termo - CRUD
type Termo struct {
	ID            int       `json:"id,omitempty"`
	Nome          string    `json:"nome,omitempty"`
	DataCadastro  string    `json:"datacadastro,omitempty"`
	DataAvaliacao string    `json:"dataavaliacao,omitempty"`
	Avaliacao     Avaliacao `json:"avaliacao,omitempty"`
	Gestor        Gestor    `json:"gestor,omitempty"`
	Empresa       Empresa   `json:"empresa,omitempty"`
	Message       controler.Message
}

//TermoList struct consulta
type TermoList struct {
	Termo    []Termo `json:"Termo"`
	Contador int     `json:"contador"`
	Message  controler.Message
}

var termos []Termo

func insertORUpTermo(obj Termo) error {

	var err error
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update Termo set nome = $1, datacadastro = $2, dataavaliacao= $3, id_avaliacao = $4, id_gestor = $5, id_empresa = $6 where id =$7")
		stmt.Exec(obj.Nome, obj.DataCadastro, obj.DataAvaliacao, obj.Avaliacao.ID, obj.Gestor.ID, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into Termo(nome, datacadastro, dataavaliacao, id_avaliacao id_gestor, id_empresa) values($1,$2,$3,$4,$5,$6) RETURNING id")
		stmt.Exec(obj.Nome, obj.DataCadastro, obj.DataAvaliacao, obj.Avaliacao.ID, obj.Gestor.ID, obj.Empresa.ID)

		if err != nil {
			log.Fatal("Cannot run insert statement", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	}
	return err
}

//CreateTermo contas
func CreateTermo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Termo

	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	err = insertORUpTermo(obj)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//DeleteTermo Termo
func DeleteTermo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Termo

	obj.ID = controler.StringForInt(vars["id"])

	if obj.ID == 0 {
		fmt.Println("Não foi possível encontrar a Etapa da Obra")
	}

	err := deletarTermo(obj.ID)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Remocao realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}
func deletarTermo(id int) error {
	stmt, err := persistencia.DB.Prepare("delete from termo where id = $1")
	stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	return err
}

//GetTermo Termo
func GetTermo(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Termo
	var err error
	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Nao foi possivel encontrar o item")
	}
	obj, err = getTermoID(obj.ID)

	json.NewEncoder(w).Encode(&obj)
}

func getTermoID(ID int) (Termo, error) {

	rows, err := persistencia.DB.Query("select * from Termo where id = $1", ID)
	if err != nil {
		log.Fatal("Erro getTermoID", err)
	}
	var ps []Termo
	for rows.Next() {
		var um Termo
		rows.Scan(&um.ID, &um.Nome, &um.DataCadastro, &um.DataAvaliacao, &um.Avaliacao.ID, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Termo{ID: um.ID, Nome: um.Nome, DataCadastro: um.DataCadastro, DataAvaliacao: um.DataAvaliacao, Avaliacao: um.Avaliacao, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	var obj Termo
	for _, numero := range ps {
		obj = numero
	}
	return obj, err
}

//GetTermoAll Termo todos
func GetTermoAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj Termo
	var objList TermoList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.Termo, erro = getTermoAll(obj.Gestor.ID)
	objList.Contador = len(objList.Termo)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 304
	}

	json.NewEncoder(w).Encode(&objList)

}
func getTermoAll(id int) ([]Termo, error) {

	rows, err := persistencia.DB.Query("SELECT id, nome, datacadastro, dataavaliacao, id_avaliacao, id_gestor, id_empresa FROM Termo where id_empresa = $1 ORDER BY datacadastro DESC ", id)
	if err != nil {
		log.Fatal("Erro getTermoAll", err)
	}
	var ps []Termo
	for rows.Next() {
		var um Termo
		rows.Scan(&um.ID, &um.Nome, &um.DataCadastro, &um.DataAvaliacao, &um.Avaliacao.ID, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Termo{ID: um.ID, Nome: um.Nome, DataCadastro: um.DataCadastro, DataAvaliacao: um.DataAvaliacao, Avaliacao: um.Avaliacao, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

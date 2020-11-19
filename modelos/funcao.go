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

//Funcao - CRUD
type Funcao struct {
	ID           int     `json:"id,omitempty"`
	Nome         string  `json:"nome,omitempty"`
	DataCadastro string  `json:"datacadastro,omitempty"`
	Gestor       Gestor  `json:"gestor,omitempty"`
	Empresa      Empresa `json:"empresa,omitempty"`
	Message      controler.Message
}

//FuncaoList struct consulta
type FuncaoList struct {
	Funcao   []Funcao `json:"funcao"`
	Contador int      `json:"contador"`
	Message  controler.Message
}

var funcaos []Funcao

func insertORUpFuncao(obj Funcao) error {

	var err error
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update Funcao set nome = $1, datacadastro = $2, id_gestor = $3, id_empresa = $4 where id =$5")
		stmt.Exec(obj.Nome, obj.DataCadastro, obj.Gestor.ID, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into Funcao(nome, datacadastro, id_gestor, id_empresa) values($1,$2,$3,$4) RETURNING id")
		stmt.Exec(obj.Nome, obj.DataCadastro, obj.Gestor.ID, obj.Empresa.ID)

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

//CreateFuncao contas
func CreateFuncao(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Funcao

	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	err = insertORUpFuncao(obj)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//DeleteFuncao Funcao
func DeleteFuncao(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Funcao

	obj.ID = controler.StringForInt(vars["id"])

	if obj.ID == 0 {
		fmt.Println("Não foi possível encontrar a Etapa da Obra")
	}

	err := deletarFuncao(obj.ID)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Remocao realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}
func deletarFuncao(id int) error {
	stmt, err := persistencia.DB.Prepare("delete from Funcao where id = $1")
	stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	return err
}

//GetFuncao Funcao
func GetFuncao(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Funcao
	var err error
	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Nao foi possivel encontrar o item")
	}
	obj, err = getFuncaoID(obj.ID)

	json.NewEncoder(w).Encode(&obj)
}

func getFuncaoID(ID int) (Funcao, error) {

	rows, err := persistencia.DB.Query("select * from Funcao where id = $1", ID)
	if err != nil {
		log.Fatal("Erro getFuncaoID", err)
	}
	var ps []Funcao
	for rows.Next() {
		var um Funcao
		rows.Scan(&um.ID, &um.Nome, &um.DataCadastro, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Funcao{ID: um.ID, Nome: um.Nome, DataCadastro: um.DataCadastro, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	var obj Funcao
	for _, numero := range ps {
		obj = numero
	}
	return obj, err
}

//GetFuncaoAll Funcao todos
func GetFuncaoAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj Funcao
	var objList FuncaoList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.Funcao, erro = getFuncaoAll(obj.Empresa.ID)
	objList.Contador = len(objList.Funcao)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 304
	}

	json.NewEncoder(w).Encode(&objList)

}
func getFuncaoAll(id int) ([]Funcao, error) {

	rows, err := persistencia.DB.Query("SELECT id, nome, datacadastro, id_gestor, id_empresa FROM Funcao where id_empresa = $1 ORDER BY datacadastro DESC ", id)
	if err != nil {
		log.Fatal("Erro getFuncaoAll", err)
	}
	var ps []Funcao
	for rows.Next() {
		var um Funcao
		rows.Scan(&um.ID, &um.Nome, &um.DataCadastro, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Funcao{ID: um.ID, Nome: um.Nome, DataCadastro: um.DataCadastro, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

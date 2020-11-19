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

//Grauinstrucao - CRUD
type Grauinstrucao struct {
	ID           int     `json:"id,omitempty"`
	Nome         string  `json:"nome,omitempty"`
	DataCadastro string  `json:"datacadastro,omitempty"`
	Gestor       Gestor  `json:"gestor,omitempty"`
	Empresa      Empresa `json:"empresa,omitempty"`
	Message      controler.Message
}

//GrauinstrucaoList struct consulta
type GrauinstrucaoList struct {
	Grauinstrucao []Grauinstrucao `json:"grauinstrucao"`
	Contador      int             `json:"contador"`
	Message       controler.Message
}

var grauinstrucao []Grauinstrucao

func insertORUpGrauinstrucao(obj Grauinstrucao) error {

	var err error
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update grauinstrucao set nome = $1, datacadastro = $2, id_gestor = $3, id_empresa = $4 where id =$5")
		stmt.Exec(obj.Nome, obj.DataCadastro, obj.Gestor.ID, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into grauinstrucao(nome, datacadastro, id_gestor, id_empresa) values($1,$2,$3,$4) RETURNING id")
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

//CreateGrauinstrucao contas
func CreateGrauinstrucao(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Grauinstrucao

	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	err = insertORUpGrauinstrucao(obj)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//DeleteGrauinstrucao Grauinstrucao
func DeleteGrauinstrucao(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Grauinstrucao

	obj.ID = controler.StringForInt(vars["id"])

	if obj.ID == 0 {
		fmt.Println("Não foi possível encontrar a Etapa da Obra")
	}

	err := deletarGrauinstrucao(obj.ID)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Remocao realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}
func deletarGrauinstrucao(id int) error {
	stmt, err := persistencia.DB.Prepare("delete from Grauinstrucao where id = $1")
	stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	return err
}

//GetGrauinstrucao Grauinstrucao
func GetGrauinstrucao(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Grauinstrucao
	var err error
	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Nao foi possivel encontrar o item")
	}
	obj, err = getGrauinstrucaoID(obj.ID)

	json.NewEncoder(w).Encode(&obj)
}

func getGrauinstrucaoID(ID int) (Grauinstrucao, error) {

	rows, err := persistencia.DB.Query("select * from grauinstrucao where id = $1", ID)
	if err != nil {
		log.Fatal("Erro getGrauinstrucaoID", err)
	}
	var ps []Grauinstrucao
	for rows.Next() {
		var um Grauinstrucao
		rows.Scan(&um.ID, &um.Nome, &um.DataCadastro, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Grauinstrucao{ID: um.ID, Nome: um.Nome, DataCadastro: um.DataCadastro, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	var obj Grauinstrucao
	for _, numero := range ps {
		obj = numero
	}
	return obj, err
}

//GetGrauinstrucaoAll Grauinstrucao todos
func GetGrauinstrucaoAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj Grauinstrucao
	var objList GrauinstrucaoList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.Grauinstrucao, erro = getGrauinstrucaoAll(obj.Empresa.ID)
	objList.Contador = len(objList.Grauinstrucao)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 304
	}

	json.NewEncoder(w).Encode(&objList)

}
func getGrauinstrucaoAll(id int) ([]Grauinstrucao, error) {

	rows, err := persistencia.DB.Query("SELECT id, nome, datacadastro, id_gestor, id_empresa FROM Grauinstrucao where id_empresa = $1 ORDER BY datacadastro DESC ", id)
	if err != nil {
		log.Fatal("Erro getGrauinstrucaoAll", err)
	}
	var ps []Grauinstrucao
	for rows.Next() {
		var um Grauinstrucao
		rows.Scan(&um.ID, &um.Nome, &um.DataCadastro, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Grauinstrucao{ID: um.ID, Nome: um.Nome, DataCadastro: um.DataCadastro, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

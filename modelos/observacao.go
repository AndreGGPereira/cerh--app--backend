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

//Observacao - CRUD
type Observacao struct {
	ID           int       `json:"id,omitempty"`
	DataCadastro string    `json:"datacadastro,omitempty"`
	Obs          string    `json:"obs,omitempty"`
	Melhorias    string    `json:"melhorias,omitempty"`
	Avaliacao    Avaliacao `json:"avaliacao,omitempty"`
	Gestor       Gestor    `json:"gestor,omitempty"`
	Empresa      Empresa   `json:"empresa,omitempty"`
	Message      controler.Message
}

//ObservacaoList struct consulta
type ObservacaoList struct {
	Observacao []Observacao `json:"Observacao"`
	Contador   int          `json:"contador"`
	Message    controler.Message
}

var observacaos []Observacao

func insertORUpObservacao(obj Observacao) error {

	var err error
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update Observacao set datacadastro = $1, obs = $2, melhorias = $3, id_avaliacao = $4, id_gestor = $5, id_empresa = $6 where id =$7")
		stmt.Exec(obj.DataCadastro, obj.Obs, obj.Melhorias, obj.Avaliacao, obj.Gestor.ID, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into Observacao(datacadastro, obs, melhorias, id_avaliacao, id_gestor, id_empresa) values($1,$2,$3,$4,$5,$6) RETURNING id")
		stmt.Exec(obj.DataCadastro, obj.Obs, obj.Melhorias, obj.Avaliacao, obj.Gestor.ID, obj.Empresa.ID)

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

//CreateObservacao contas
func CreateObservacao(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Observacao

	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	err = insertORUpObservacao(obj)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//DeleteObservacao Observacao
func DeleteObservacao(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Observacao

	obj.ID = controler.StringForInt(vars["id"])

	if obj.ID == 0 {
		fmt.Println("Não foi possível encontrar a Etapa da Obra")
	}

	err := deletarObservacao(obj.ID)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Remocao realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}
func deletarObservacao(id int) error {
	stmt, err := persistencia.DB.Prepare("delete from Observacao where id = $1")
	stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	return err
}

//GetObservacao Observacao
func GetObservacao(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Observacao
	var err error
	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Nao foi possivel encontrar o item")
	}
	obj, err = getObservacaoID(obj.ID)

	json.NewEncoder(w).Encode(&obj)
}

func getObservacaoID(ID int) (Observacao, error) {

	rows, err := persistencia.DB.Query("select * from Observacao where id = $1", ID)
	if err != nil {
		log.Fatal("Erro getObservacaoID", err)
	}
	var ps []Observacao
	for rows.Next() {
		var um Observacao
		rows.Scan(&um.ID, &um.DataCadastro, &um.Obs, &um.Melhorias, &um.Avaliacao.ID, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Observacao{ID: um.ID, DataCadastro: um.DataCadastro, Obs: um.Obs, Melhorias: um.Melhorias, Avaliacao: um.Avaliacao, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	var obj Observacao
	for _, numero := range ps {
		obj = numero
	}
	return obj, err
}

//GetObservacaoAll Observacao todos
func GetObservacaoAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj Observacao
	var objList ObservacaoList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.Observacao, erro = getObservacaoAll(obj.Empresa.ID)
	objList.Contador = len(objList.Observacao)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 304
	}

	json.NewEncoder(w).Encode(&objList)

}
func getObservacaoAll(id int) ([]Observacao, error) {

	rows, err := persistencia.DB.Query("SELECT id, datacadastra, obs, melhorias, id_avaliacao, id_gestor, id_empresa FROM Observacao where id_empresa = $1 ORDER BY datacadastro DESC ", id)
	if err != nil {
		log.Fatal("Erro getObservacaoAll", err)
	}
	var ps []Observacao
	for rows.Next() {
		var um Observacao
		rows.Scan(&um.ID, &um.DataCadastro, &um.Obs, &um.Melhorias, &um.Avaliacao.ID, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Observacao{ID: um.ID, DataCadastro: um.DataCadastro, Obs: um.Obs, Melhorias: um.Melhorias, Avaliacao: um.Avaliacao, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

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

//TipoCompetencia - CRUD
type TipoCompetencia struct {
	ID           int     `json:"id,omitempty"`
	Nome         string  `json:"nome,omitempty"`
	DataCadastro string  `json:"datacadastro,omitempty"`
	Abreviacao   string  `json:"abreviacao,omitempty"`
	Descricao    string  `json:"descricao,omitempty"`
	Gestor       Gestor  `json:"gestor,omitempty"`
	Empresa      Empresa `json:"empresa,omitempty"`
	Message      controler.Message
}

//TipoCompetenciaList struct consulta
type TipoCompetenciaList struct {
	TipoCompetencia []TipoCompetencia `json:"tipocompetencia"`
	Contador        int               `json:"contador"`
	Message         controler.Message
}

var tipocompetencias []TipoCompetencia

func insertORUpTipoCompetencia(obj TipoCompetencia) error {

	var err error
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update TipoCompetencia set nome = $1, datacadastro = $2, descricao= $3, abreviacao= $4, id_gestor = $5, id_empresa = $6 where id =$7")
		stmt.Exec(obj.Nome, obj.DataCadastro, obj.Descricao, obj.Abreviacao, obj.Gestor.ID, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into TipoCompetencia(nome, datacadastro, descricao, abreviacao, id_gestor, id_empresa) values($1,$2,$3,$4,$5,$6) RETURNING id")
		stmt.Exec(obj.Nome, obj.DataCadastro, obj.Descricao, obj.Abreviacao, obj.Gestor.ID, obj.Empresa.ID)

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

//CreateTipoCompetencia contas
func CreateTipoCompetencia(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj TipoCompetencia

	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	err = insertORUpTipoCompetencia(obj)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//DeleteTipoCompetencia tipo competencia
func DeleteTipoCompetencia(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj TipoCompetencia

	obj.ID = controler.StringForInt(vars["id"])

	if obj.ID == 0 {
		fmt.Println("Não foi possível encontrar a Etapa da Obra")
	}

	err := deletarTipoCompetencia(obj.ID)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Remocao realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}
func deletarTipoCompetencia(id int) error {
	stmt, err := persistencia.DB.Prepare("delete from TipoCompetencia where id = $1")
	stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	return err
}

//GetTipoCompetencia TipoCompetencia
func GetTipoCompetencia(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj TipoCompetencia
	var err error
	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Nao foi possivel encontrar o item")
	}
	obj, err = getTipoCompetenciaID(obj.ID)

	json.NewEncoder(w).Encode(&obj)
}

func getTipoCompetenciaID(ID int) (TipoCompetencia, error) {

	rows, err := persistencia.DB.Query("select * from TipoCompetencia where id = $1", ID)
	if err != nil {
		log.Fatal("Erro TipoCompetencia", err)
	}
	var ps []TipoCompetencia
	for rows.Next() {
		var um TipoCompetencia
		rows.Scan(&um.ID, &um.Nome, &um.DataCadastro, &um.Descricao, &um.Abreviacao, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, TipoCompetencia{ID: um.ID, Nome: um.Nome, DataCadastro: um.DataCadastro, Descricao: um.Descricao, Abreviacao: um.Abreviacao, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	var obj TipoCompetencia
	for _, numero := range ps {
		obj = numero
	}
	return obj, err
}

//GetTipoCompetenciaAll TipoCompetencia todos
func GetTipoCompetenciaAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj TipoCompetencia
	var objList TipoCompetenciaList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.TipoCompetencia, erro = getTipoCompetenciaAll(obj.Empresa.ID)
	objList.Contador = len(objList.TipoCompetencia)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 302
	}

	json.NewEncoder(w).Encode(&objList)

}
func getTipoCompetenciaAll(id int) ([]TipoCompetencia, error) {

	rows, err := persistencia.DB.Query("SELECT id, nome, datacadastro, descricao, abreviacao, id_gestor, id_empresa FROM tipocompetencia where id_empresa = $1 ORDER BY datacadastro DESC ", id)
	if err != nil {
		log.Fatal("Erro getTipoCompetenciaAll", err)
	}
	var ps []TipoCompetencia
	for rows.Next() {
		var um TipoCompetencia
		rows.Scan(&um.ID, &um.Nome, &um.DataCadastro, &um.Descricao, &um.Abreviacao, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, TipoCompetencia{ID: um.ID, Nome: um.Nome, DataCadastro: um.DataCadastro, Descricao: um.Descricao, Abreviacao: um.Abreviacao, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

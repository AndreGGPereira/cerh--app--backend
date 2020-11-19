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

//CompetenciaPa - CRUD
type CompetenciaPa struct {
	ID              int             `json:"id,omitempty"`
	DataCadastro    string          `json:"datacadastro,omitempty"`
	Titulo          string          `json:"titulo,omitempty"`
	Descricao       string          `json:"descricao,omitempty"`
	Ativo           bool            `json:"ativo,omitempty"`
	TipoCompetencia TipoCompetencia `json:"tipocompetencia,omitempty"`
	Gestor          Gestor          `json:"gestor,omitempty"`
	Empresa         Empresa         `json:"empresa,omitempty"`
	Message         controler.Message
}

//CompetenciaPaList struct consulta
type CompetenciaPaList struct {
	CompetenciaPa []CompetenciaPa   `json:"competenciapa,omitempty"`
	Contador      int               `json:"contador,omitempty"`
	Message       controler.Message `json:"message,omitempty"`
}

var competenciaPas []CompetenciaPa

func insertORUpCompetenciaPa(obj CompetenciaPa) error {

	var err error
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update CompetenciaPa set titulo = $1, datacadastro = $2, descricao= $3, ativo= $4, id_tipocompetencia = $5, id_gestor = $6, id_empresa = $7 where id =$8")
		stmt.Exec(obj.Titulo, obj.DataCadastro, obj.Descricao, obj.Ativo, obj.TipoCompetencia.ID, obj.Gestor.ID, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into CompetenciaPa(titulo, datacadastro, descricao, ativo, id_tipocompetencia, id_gestor, id_empresa) values($1,$2,$3,$4,$5,$6,$7) RETURNING id")
		stmt.Exec(obj.Titulo, obj.DataCadastro, obj.Descricao, obj.Ativo, obj.TipoCompetencia.ID, obj.Gestor.ID, obj.Empresa.ID)

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

//CreateCompetenciaPa contas
func CreateCompetenciaPa(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj CompetenciaPa
	err := decoder.Decode(&obj)

	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	err = insertORUpCompetenciaPa(obj)

	if err == nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//CreateCompetenciaPaAll cadastrar lista
func CreateCompetenciaPaAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj CompetenciaPaList

	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	for _, dado := range obj.CompetenciaPa {
		var CompetenciaPa CompetenciaPa

		CompetenciaPa.Ativo = dado.Ativo
		CompetenciaPa.DataCadastro = controler.PegarDataAtualStringNew()
		CompetenciaPa.Descricao = dado.Descricao
		CompetenciaPa.Empresa.ID, CompetenciaPa.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
		CompetenciaPa.TipoCompetencia = dado.TipoCompetencia
		CompetenciaPa.Titulo = dado.Titulo
		err = insertORUpCompetenciaPa(CompetenciaPa)
	}

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//DeleteCompetenciaPa CompetenciaPa
func DeleteCompetenciaPa(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj CompetenciaPa

	obj.ID = controler.StringForInt(vars["id"])

	if obj.ID == 0 {
		fmt.Println("Não foi possível encontrar a Etapa da Obra")
	}

	err := deletarCompetenciaPa(obj.ID)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Remocao realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}
func deletarCompetenciaPa(id int) error {
	stmt, err := persistencia.DB.Prepare("delete from CompetenciaPa where id = $1")
	stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	return err
}

//GetCompetenciaPa CompetenciaPa
func GetCompetenciaPa(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj CompetenciaPa
	var err error
	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Nao foi possivel encontrar o item")
	}
	obj, err = getCompetenciaPaID(obj.ID)

	json.NewEncoder(w).Encode(&obj)
}

func getCompetenciaPaID(ID int) (CompetenciaPa, error) {

	rows, err := persistencia.DB.Query("select * from CompetenciaPa where id = $1", ID)
	if err != nil {
		log.Fatal("Erro getCompetenciaPaID", err)
	}
	var ps []CompetenciaPa
	for rows.Next() {
		var um CompetenciaPa
		rows.Scan(&um.ID, &um.DataCadastro, &um.Titulo, &um.Descricao, &um.Ativo, &um.TipoCompetencia.ID, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, CompetenciaPa{ID: um.ID, DataCadastro: um.DataCadastro, Titulo: um.Titulo, Descricao: um.Descricao, Ativo: um.Ativo, TipoCompetencia: um.TipoCompetencia, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	var obj CompetenciaPa
	for _, numero := range ps {
		obj = numero
	}
	return obj, err
}

//GetCompetenciaPaAll CompetenciaPa todos
func GetCompetenciaPaAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj CompetenciaPa
	var objList CompetenciaPaList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.CompetenciaPa, erro = getCompetenciaPaAll(obj.Empresa.ID)
	objList.Contador = len(objList.CompetenciaPa)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 202
	}

	json.NewEncoder(w).Encode(&objList)

}

func getCompetenciaPaAll(id int) ([]CompetenciaPa, error) {

	rows, err := persistencia.DB.Query("SELECT cp.id, cp.datacadastro, cp.titulo, cp.descricao, cp.ativo, cp.id_tipocompetencia, cp.id_gestor, cp.id_empresa, tp.nome FROM CompetenciaPa AS cp INNER JOIN tipocompetencia as tp ON cp.id_tipocompetencia = tp.id where cp.id_empresa = $1 ORDER BY cp.titulo ASC ", id)
	if err != nil {
		log.Fatal("Erro getCompetenciaPaAll", err)
	}
	var ps []CompetenciaPa
	for rows.Next() {
		var um CompetenciaPa
		rows.Scan(&um.ID, &um.DataCadastro, &um.Titulo, &um.Descricao, &um.Ativo, &um.TipoCompetencia.ID, &um.Gestor.ID, &um.Empresa.ID, &um.TipoCompetencia.Nome)
		ps = append(ps, CompetenciaPa{ID: um.ID, DataCadastro: um.DataCadastro, Titulo: um.Titulo, Descricao: um.Descricao, Ativo: um.Ativo, TipoCompetencia: um.TipoCompetencia, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

//GetCompetenciaPaTECAll CompetenciaPa todos
func GetCompetenciaPaTECAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj CompetenciaPa
	var objList CompetenciaPaList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.CompetenciaPa, erro = getCompetenciaPaTECAll(obj.Empresa.ID)
	objList.Contador = len(objList.CompetenciaPa)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 202
	}

	json.NewEncoder(w).Encode(&objList)

}

func getCompetenciaPaTECAll(id int) ([]CompetenciaPa, error) {

	rows, err := persistencia.DB.Query("SELECT cp.id, cp.datacadastro, cp.titulo, cp.descricao, cp.ativo, cp.id_tipocompetencia, cp.id_gestor, cp.id_empresa, tp.nome FROM CompetenciaPa AS cp INNER JOIN tipocompetencia as tp ON cp.id_tipocompetencia = tp.id where cp.id_empresa = $1 and tp.nome = 'TÉCNICA' ORDER BY cp.titulo ASC", id)
	if err != nil {
		log.Fatal("Erro getCompetenciaPaAll", err)
	}
	var ps []CompetenciaPa
	for rows.Next() {
		var um CompetenciaPa
		rows.Scan(&um.ID, &um.DataCadastro, &um.Titulo, &um.Descricao, &um.Ativo, &um.TipoCompetencia.ID, &um.Gestor.ID, &um.Empresa.ID, &um.TipoCompetencia.Nome)
		ps = append(ps, CompetenciaPa{ID: um.ID, DataCadastro: um.DataCadastro, Titulo: um.Titulo, Descricao: um.Descricao, Ativo: um.Ativo, TipoCompetencia: um.TipoCompetencia, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

//GetCompetenciaPaCOMAll CompetenciaPa todos
func GetCompetenciaPaCOMAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj CompetenciaPa
	var objList CompetenciaPaList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.CompetenciaPa, erro = getCompetenciaPaCOMAll(obj.Empresa.ID)
	objList.Contador = len(objList.CompetenciaPa)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 202
	}

	json.NewEncoder(w).Encode(&objList)

}

func getCompetenciaPaCOMAll(id int) ([]CompetenciaPa, error) {

	rows, err := persistencia.DB.Query("SELECT cp.id, cp.datacadastro, cp.titulo, cp.descricao, cp.ativo, cp.id_tipocompetencia, cp.id_gestor, cp.id_empresa, tp.nome FROM CompetenciaPa AS cp INNER JOIN tipocompetencia as tp ON cp.id_tipocompetencia = tp.id where cp.id_empresa = $1 and tp.nome = 'COMPORTAMENTAL' ORDER BY cp.titulo ASC ", id)
	if err != nil {
		log.Fatal("Erro getCompetenciaPaAll", err)
	}
	var ps []CompetenciaPa
	for rows.Next() {
		var um CompetenciaPa
		rows.Scan(&um.ID, &um.DataCadastro, &um.Titulo, &um.Descricao, &um.Ativo, &um.TipoCompetencia.ID, &um.Gestor.ID, &um.Empresa.ID, &um.TipoCompetencia.Nome)
		ps = append(ps, CompetenciaPa{ID: um.ID, DataCadastro: um.DataCadastro, Titulo: um.Titulo, Descricao: um.Descricao, Ativo: um.Ativo, TipoCompetencia: um.TipoCompetencia, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

//GetCompetenciaPaORGAll CompetenciaPa todos
func GetCompetenciaPaORGAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj CompetenciaPa
	var objList CompetenciaPaList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.CompetenciaPa, erro = getCompetenciaPaORGAll(obj.Empresa.ID)
	objList.Contador = len(objList.CompetenciaPa)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 202
	}

	json.NewEncoder(w).Encode(&objList)

}

func getCompetenciaPaORGAll(id int) ([]CompetenciaPa, error) {

	rows, err := persistencia.DB.Query("SELECT cp.id, cp.datacadastro, cp.titulo, cp.descricao, cp.ativo, cp.id_tipocompetencia, cp.id_gestor, cp.id_empresa, tp.nome FROM CompetenciaPa AS cp INNER JOIN tipocompetencia as tp ON cp.id_tipocompetencia = tp.id where cp.id_empresa = $1 and tp.nome = 'ORGANIZACIONAL' ORDER BY cp.titulo ASC ", id)
	if err != nil {
		log.Fatal("Erro getCompetenciaPaAll", err)
	}
	var ps []CompetenciaPa
	for rows.Next() {
		var um CompetenciaPa
		rows.Scan(&um.ID, &um.DataCadastro, &um.Titulo, &um.Descricao, &um.Ativo, &um.TipoCompetencia.ID, &um.Gestor.ID, &um.Empresa.ID, &um.TipoCompetencia.Nome)
		ps = append(ps, CompetenciaPa{ID: um.ID, DataCadastro: um.DataCadastro, Titulo: um.Titulo, Descricao: um.Descricao, Ativo: um.Ativo, TipoCompetencia: um.TipoCompetencia, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

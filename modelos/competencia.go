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

//Competencia - CRUD
type Competencia struct {
	ID              int             `json:"id,omitempty"`
	Titulo          string          `json:"titulo,omitempty"`
	Abreviacao      string          `json:"abreviacao,omitempty"`
	Descricao       string          `json:"descricao,omitempty"`
	DataCadastro    string          `json:"datacadastro,omitempty"`
	Ativo           bool            `json:"ativo,omitempty"`
	TipoCompetencia TipoCompetencia `json:"tipocompetencia,omitempty"`
	Cargo           Cargo           `json:"cargo,omitempty"`
	Gestor          Gestor          `json:"gestor,omitempty"`
	Empresa         Empresa         `json:"empresa,omitempty"`
	Message         controler.Message
}

//CompetenciaList struct consulta
type CompetenciaList struct {
	Competencia []Competencia     `json:"competencia,omitempty"`
	Contador    int               `json:"contador,omitempty"`
	Message     controler.Message `json:"message,omitempty"`
}

var competencias []Competencia

func insertORUpCompetencia(obj Competencia) error {

	var err error
	if obj.ID != 0 {
		fmt.Println("Entrou aqui insertORUpCompetencia line 45", obj)
		stmt, err := persistencia.DB.Prepare("update competencia set titulo = $1, abreviacao= $2, descricao= $3, datacadastro= $4, ativo= $5, id_tipocompetencia = $6, id_cargo = $7, id_gestor = $8, id_empresa = $9 where id =$10")
		stmt.Exec(obj.Titulo, obj.Abreviacao, obj.Descricao, obj.DataCadastro, obj.Ativo, obj.TipoCompetencia.ID, obj.Cargo.ID, obj.Gestor.ID, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		fmt.Println("Entrou aqui insertORUpCompetencia line 59", obj)
		stmt, err := persistencia.DB.Prepare("insert into competencia(titulo, abreviacao, descricao, datacadastro, ativo, id_tipocompetencia, id_cargo, id_gestor, id_empresa) values($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id")
		stmt.Exec(obj.Titulo, obj.Abreviacao, obj.Descricao, obj.DataCadastro, obj.Ativo, obj.TipoCompetencia.ID, obj.Cargo.ID, obj.Gestor.ID, obj.Empresa.ID)

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

//CreateCompetencia contas
func CreateCompetencia(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Competencia

	err := decoder.Decode(&obj)

	fmt.Println("Entrou aqui CreateCompetencia line 82", obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	err = insertORUpCompetencia(obj)

	if err == nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//CreateCompetenciaAll cadastrar lista
func CreateCompetenciaAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj CompetenciaList

	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	for _, dado := range obj.Competencia {
		var competencia Competencia

		competencia.Ativo = dado.Ativo
		competencia.Cargo = dado.Cargo
		competencia.DataCadastro = controler.PegarDataAtualStringNew()
		competencia.Descricao = dado.Descricao
		competencia.Empresa.ID, competencia.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
		competencia.TipoCompetencia = dado.TipoCompetencia
		competencia.Titulo = dado.Titulo

		fmt.Println("Entrou aqui CreateCompetencia line 127", competencia.Cargo.ID)
		fmt.Println("Entrou aqui CreateCompetencia line 127", competencia.TipoCompetencia.ID)

		err = insertORUpCompetencia(competencia)
	}

	fmt.Println("Dados do erro a insercao", err)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//DeleteCompetencia Competencia
func DeleteCompetencia(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Competencia

	obj.ID = controler.StringForInt(vars["id"])

	if obj.ID == 0 {
		fmt.Println("Não foi possível encontrar a Etapa da Obra")
	}

	err := deletarCompetencia(obj.ID)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Remocao realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}
func deletarCompetencia(id int) error {
	stmt, err := persistencia.DB.Prepare("delete from competencia where id = $1")
	stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	return err
}

//GetCompetencia Competencia
func GetCompetencia(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Competencia
	var err error
	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Nao foi possivel encontrar o item")
	}
	obj, err = getCompetenciaID(obj.ID)

	json.NewEncoder(w).Encode(&obj)
}

func getCompetenciaID(ID int) (Competencia, error) {

	rows, err := persistencia.DB.Query("select * from Competencia where id = $1", ID)
	if err != nil {
		log.Fatal("Erro getCompetenciaID", err)
	}
	var ps []Competencia
	for rows.Next() {
		var um Competencia
		rows.Scan(&um.ID, &um.Titulo, &um.Abreviacao, &um.Descricao, &um.DataCadastro, &um.Ativo, &um.TipoCompetencia.ID, &um.Cargo.ID, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Competencia{ID: um.ID, Titulo: um.Titulo, Abreviacao: um.Abreviacao, Descricao: um.Descricao, DataCadastro: um.DataCadastro, Ativo: um.Ativo, TipoCompetencia: um.TipoCompetencia, Cargo: um.Cargo, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	var obj Competencia
	for _, numero := range ps {
		obj = numero
	}
	return obj, err
}

//GetCompetenciaForJob Competencia
func GetCompetenciaForJob(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Competencia
	var objList CompetenciaList
	var err error

	obj.Cargo.ID, err = strconv.Atoi(vars["id"])

	objList.Competencia, err = getCompetenciaForJob(obj.Cargo.ID)

	fmt.Println(" objList.Competencia ", objList.Competencia)
	fmt.Println(" objList.(obj.Cargo.ID ", obj.Cargo.ID)
	objList.Contador = len(objList.Competencia)
	fmt.Println(" Dados do Competencia ", objList.Contador)

	if err != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 202
	}

	json.NewEncoder(w).Encode(&objList)
}
func getCompetenciaForJob(id int) ([]Competencia, error) {
	//rows, _ := persistencia.DB.Query("select eo.id, eo.nome, eo.descricao, eo.datacadastro, eo.datainicioetapa, eo.datafimetapa, eo.ativa, eo.id_obra, eo.id_usuario, eo.id_empresa, eo.id_etapaobratipo, eo.status, et.nome, u.nome from EtapaObra AS eo INNER JOIN etapaobratipo as et ON eo.id_etapaobratipo = et.id INNER JOIN usuario as u ON eo.id_usuario = u.id where id_obra = $1 ORDER BY  eo.datacadastro DESC", ID)
	rows, err := persistencia.DB.Query("SELECT c.id, c.titulo, c.abreviacao, c.descricao, c.datacadastro, c.ativo, c.id_tipocompetencia, c.id_cargo, c.id_gestor, c.id_empresa, ge.nome FROM Competencia AS c INNER JOIN gestor as ge ON c.id_gestor = ge.id where c.id_cargo = $1 ORDER BY c.titulo ASC ", id)

	if err != nil {
		fmt.Println(err)
	}

	var ps []Competencia
	for rows.Next() {
		var um Competencia
		rows.Scan(&um.ID, &um.Titulo, &um.Abreviacao, &um.Descricao, &um.DataCadastro, &um.Ativo, &um.TipoCompetencia.ID, &um.Cargo.ID, &um.Gestor.ID, &um.Empresa.ID, &um.Gestor.Nome)
		ps = append(ps, Competencia{ID: um.ID, Titulo: um.Titulo, Abreviacao: um.Abreviacao, Descricao: um.Descricao, DataCadastro: um.DataCadastro, Ativo: um.Ativo, TipoCompetencia: um.TipoCompetencia, Cargo: um.Cargo, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

//GetCompetenciaAll Competencia todos
func GetCompetenciaAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj Competencia
	var objList CompetenciaList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.Competencia, erro = getCompetenciaAll(obj.Empresa.ID)
	objList.Contador = len(objList.Competencia)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 202
	}

	json.NewEncoder(w).Encode(&objList)

}
func getCompetenciaAll(id int) ([]Competencia, error) {

	rows, err := persistencia.DB.Query("SELECT id, titulo, datacadastro, ativo, abreviacao, descricao, id_tipocompetencia, id_cargo, id_gestor, id_empresa FROM Competencia where id_empresa = $1 ORDER BY titulo ASC ", id)
	if err != nil {
		log.Fatal("Erro getCompetenciaAll", err)
	}
	var ps []Competencia
	for rows.Next() {
		var um Competencia
		rows.Scan(&um.ID, &um.Titulo, &um.Abreviacao, &um.Descricao, &um.DataCadastro, &um.Ativo, &um.TipoCompetencia.ID, &um.Cargo.ID, &um.Gestor.ID, &um.Empresa.ID, &um.Gestor.Nome)
		ps = append(ps, Competencia{ID: um.ID, Titulo: um.Titulo, Abreviacao: um.Abreviacao, Descricao: um.Descricao, DataCadastro: um.DataCadastro, Ativo: um.Ativo, TipoCompetencia: um.TipoCompetencia, Cargo: um.Cargo, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

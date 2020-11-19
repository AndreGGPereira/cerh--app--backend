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

//Tarefas - CRUD
type Tarefas struct {
	ID           int     `json:"id,omitempty"`
	Descricao    string  `json:"descricao,omitempty"`
	DataCadastro string  `json:"datacadastro,omitempty"`
	Concluido    bool    `json:"concluido,omitempty"`
	Gestor       Gestor  `json:"gestor,omitempty"`
	Empresa      Empresa `json:"empresa,omitempty"`
}

var tarefas []Tarefas

func insertORUpTarefas(obj Tarefas) {

	//Se exister ID item deverá ser atualizado
	if obj.ID != 0 {

		fmt.Println(" Dados kkkkk", obj.DataCadastro)

		stmt, err := persistencia.DB.Prepare("update tarefas set descricao = $1, datacadastro = $2, concluido = $3, id_gestor = $4, id_empresa = $5 where id =$6")
		stmt.Exec(obj.Descricao, obj.DataCadastro, obj.Concluido, obj.Gestor.ID, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into tarefas(descricao, datacadastro, concluido, id_gestor, id_empresa) values($1,$2,$3,$4,$5) RETURNING id")
		stmt.Exec(obj.Descricao, obj.DataCadastro, obj.Concluido, obj.Gestor.ID, obj.Empresa.ID)
		// stmt, err := persistencia.DB.Prepare("insert into ofpagamento(datacadastro, medicao, status, valorpago, id_of, id_pgforma, id_usuario, id_empresa) values($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id")
		//stmt.Exec(OBJ.DataCadastro, OBJ.Medicao, OBJ.Status, OBJ.ValorPago, OBJ.OF.ID, OBJ.PGForma.ID, OBJ.Usuario.ID, OBJ.Empresa.ID)
		// stmt, err := persistencia.DB.Prepare("insert into tarefas(descricao, id_usuario, id_empresa) values($1,$2,$3) RETURNING id")
		// stmt.Exec(obj.Descricao, obj.Usuario.ID, obj.Empresa)

		if err != nil {
			log.Fatal("Cannot run insert statement", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	}

}

//CreateTarefas contas
func CreateTarefas(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Tarefas
	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	fmt.Println(" dados da Empresa ", obj.Empresa.ID)
	fmt.Println(" dados da obj.Gestor.ID  ", obj.Gestor.ID)

	// if obj.Gestor.ID == 0 {
	// 	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDUsuario(req)
	// }

	// fmt.Println(" dados da Empresa ", obj.Empresa.ID)
	// fmt.Println(" dados da obj.Gestor.ID  ", obj.Gestor.ID)

	if obj.ID == 0 {
		obj.Concluido = false
	}

	insertORUpTarefas(obj)
	json.NewEncoder(w).Encode("Tarefa cadastrado com sucesso")
}

//DeleteTarefas Tarefas
func DeleteTarefas(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Tarefas

	obj.ID = controler.StringForInt(vars["id"])

	if obj.ID == 0 {
		fmt.Println("Não foi possível encontrar a Etapa da Obra")
	}

	deletarTarefas(obj.ID)
}

func deletarTarefas(id int) {
	stmt, err := persistencia.DB.Prepare("delete from Tarefas where id = $1")
	stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
}

//GetTarefas Tarefas
func GetTarefas(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, erro := strconv.Atoi(req.FormValue("id"))

	if id == 0 || erro != nil {
		fmt.Println("Não foi possível identificar a unidade medida")
	}
	var obj Tarefas
	obj = getTarefasID(id)

	json.NewEncoder(w).Encode(&obj)

}
func getTarefasID(ID int) Tarefas {

	rows, _ := persistencia.DB.Query("select * from Tarefas where id = $1", ID)

	var ps []Tarefas
	for rows.Next() {
		var um Tarefas
		rows.Scan(&um.ID, &um.Descricao, &um.DataCadastro, &um.Concluido, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Tarefas{ID: um.ID, Descricao: um.Descricao, DataCadastro: um.DataCadastro, Concluido: um.Concluido, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	var obj Tarefas
	for _, numero := range ps {
		obj = numero
	}
	return obj
}

//GetTarefasAll Tarefas todos
func GetTarefasAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj Tarefas

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)

	fmt.Println(" Ddos Tarefas all", obj.Gestor.ID)
	tarefass, erro := getTarefasAll(obj.Gestor.ID)

	if erro != nil {
		panic(erro)
	}
	json.NewEncoder(w).Encode(&tarefass)

}

func getTarefasAll(id int) ([]Tarefas, error) {

	rows, err := persistencia.DB.Query("SELECT id, descricao, datacadastro, concluido, id_gestor, id_empresa FROM tarefas where id_gestor = $1 ORDER BY datacadastro DESC ", id)
	if err != nil {
		panic(err)
	}
	var ps []Tarefas
	for rows.Next() {
		var um Tarefas
		rows.Scan(&um.ID, &um.Descricao, &um.DataCadastro, &um.Concluido, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Tarefas{ID: um.ID, Descricao: um.Descricao, DataCadastro: um.DataCadastro, Concluido: um.Concluido, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

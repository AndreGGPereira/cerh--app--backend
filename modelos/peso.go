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

//Peso - CRUD
type Peso struct {
	ID                int     `json:"id,omitempty"`
	DataCadastro      string  `json:"datacadastro,omitempty"`
	PesoSuperior      int     `json:"pesosuperior,omitempty"`
	PesoParceiro      int     `json:"pesoparceiro,omitempty"`
	PesoSubordinado   int     `json:"pesosubordinado,omitempty"`
	PesoAutoAvaliacao int     `json:"pesoautoavaliacao,omitempty"`
	Gestor            Gestor  `json:"gestor,omitempty"`
	Empresa           Empresa `json:"empresa,omitempty"`
	Message           controler.Message
}

//PesoList struct consulta
type PesoList struct {
	Peso     []Peso `json:"Peso"`
	Contador int    `json:"contador"`
	Message  controler.Message
}

var pesos []Peso

func insertORUpPeso(obj Peso) error {

	var err error
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update Peso set datacadastro = $1, pesosuperior = $2,  pesoparceiro = $3, pesosubordinado = $4, pesoautoavaliacao = $5, id_gestor = $6, id_empresa = $7 where id =$8")
		stmt.Exec(obj.DataCadastro, obj.PesoSuperior, obj.PesoParceiro, obj.PesoSubordinado, obj.PesoAutoAvaliacao, obj.Gestor.ID, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into Peso(datacadastro, pesosuperior, pesoparceiro, pesosubordinado, pesoautoavaliacao, id_gestor, id_empresa) values($1,$2,$3,$4,$5,$6,$7,8) RETURNING id")
		stmt.Exec(obj.DataCadastro, obj.PesoSuperior, obj.PesoParceiro, obj.PesoSubordinado, obj.PesoAutoAvaliacao, obj.Gestor.ID, obj.Empresa.ID)

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

//CreatePeso contas
func CreatePeso(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Peso

	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	err = insertORUpPeso(obj)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//GetPeso Peso
func GetPeso(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Peso
	var err error
	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Nao foi possivel encontrar o item")
	}
	obj, err = getPesoID(obj.ID)

	json.NewEncoder(w).Encode(&obj)
}

func getPesoID(ID int) (Peso, error) {

	rows, err := persistencia.DB.Query("select * from Peso where id = $1", ID)
	if err != nil {
		log.Fatal("Erro getPesoID", err)
	}
	var ps []Peso
	for rows.Next() {
		var um Peso
		rows.Scan(&um.ID, &um.DataCadastro, &um.PesoSuperior, &um.PesoParceiro, &um.PesoSubordinado, &um.PesoAutoAvaliacao, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Peso{ID: um.ID, DataCadastro: um.DataCadastro, PesoSuperior: um.PesoSuperior, PesoParceiro: um.PesoParceiro, PesoSubordinado: um.PesoSubordinado, PesoAutoAvaliacao: um.PesoAutoAvaliacao, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	var obj Peso
	for _, numero := range ps {
		obj = numero
	}
	return obj, err
}

//GetPesoAll Peso todos
func GetPesoAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj Peso
	var objList PesoList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.Peso, erro = getPesoAll(obj.Empresa.ID)
	objList.Contador = len(objList.Peso)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 304
	}

	json.NewEncoder(w).Encode(&objList)

}
func getPesoAll(id int) ([]Peso, error) {

	rows, err := persistencia.DB.Query("SELECT id, datacadastro, pesosuperior, pesoparceiro, pesosubordinado, pesoautoavaliacao, id_gestor, id_empresa FROM Peso where id_empresa = $1 ORDER BY datacadastro DESC ", id)
	if err != nil {
		log.Fatal("Erro getPesoAll", err)
	}
	var ps []Peso
	for rows.Next() {
		var um Peso
		rows.Scan(&um.ID, &um.DataCadastro, &um.PesoSuperior, &um.PesoParceiro, &um.PesoSubordinado, &um.PesoAutoAvaliacao, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, Peso{ID: um.ID, DataCadastro: um.DataCadastro, PesoSuperior: um.PesoSuperior, PesoParceiro: um.PesoParceiro, PesoSubordinado: um.PesoSubordinado, PesoAutoAvaliacao: um.PesoAutoAvaliacao, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	return ps, err
}

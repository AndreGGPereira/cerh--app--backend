package modelos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/andreggpereira/sisalicerce--app--backend/controler"
	"github.com/andreggpereira/sisalicerce--app--backend/persistencia"
	"github.com/gorilla/mux"
)

//PermissaoTipo tipos de permissão do Gestor
type PermissaoTipo struct {
	ID           int     `json:"id"`
	Nome         string  `json:"nome"`
	Descricao    string  `json:"descricao"`
	DataCadastro string  `json:"datacadastro"`
	Gestor       Gestor  `json:"gestor"`
	Empresa      Empresa `json:"empresa"`
}

var permissaostipo []PermissaoTipo

//atualizar ou cadastrar
func insertORUpPermissaoTipo(obj PermissaoTipo) {

	//Se exister ID item deverá ser atualizado
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update PermissaoTipo set nome = $1,descricao = $2, datacadastro = $3, id_gestor = $4, id_empresa = $5 where id =$6")
		stmt.Exec(obj.Nome, obj.Descricao, obj.DataCadastro, obj.Gestor.ID, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			panic(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into PermissaoTipo(nome, descricao, datacadastro, id_gestor, id_empresa) values($1,$2,$3,$4,$5) RETURNING id")
		stmt.Exec(obj.Nome, obj.Descricao, obj.DataCadastro, obj.Gestor.ID, obj.Empresa.ID)

		if err != nil {
			log.Fatal("Cannot run insert statement", err)
			panic(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	}
}

//CreatePermissaoTipo cadastrar permissao
func CreatePermissaoTipo(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var obj PermissaoTipo
	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.DataCadastro = controler.PegarDataAtualStringNew()

	insertORUpPermissaoTipo(obj)
	json.NewEncoder(w).Encode("Permissão cadastrado com sucesso")
}

//DeletePermissaoTipo deletar
func DeletePermissaoTipo(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range permissoes {
		if item.ID == controler.StringForInt(params["id"]) {
			permissoes = append(permissoes[:index], permissoes[index+1:]...)
			deletarPermissao(controler.StringForInt(params["id"]))
			json.NewEncoder(w).Encode("Permissão Removido com Sucesso")
			break
		}
	}
	json.NewEncoder(w).Encode(permissoes)
}

func deletarPermissaoTipo(id int) {

	stmt, err := persistencia.DB.Prepare("delete from PermissaoTipo where id = $1")
	stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	stmt.Close()
}

//GetPermissaoTipo busca item
func GetPermissaoTipo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range permissoes {
		if item.ID == controler.StringForInt(params["id"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Permissao{})

}

func getPermissaoTipoID(ID int) PermissaoTipo {

	rows, _ := persistencia.DB.Query("select id,nome,descricao,datacadastro, id_gestor, id_empresa from PermissaoTipo where id = $1", ID)
	var ps []PermissaoTipo
	for rows.Next() {
		var um PermissaoTipo
		rows.Scan(&um.ID, &um.Descricao, &um.DataCadastro, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, PermissaoTipo{ID: um.ID, Nome: um.Nome, Descricao: um.Descricao, DataCadastro: um.DataCadastro, Gestor: um.Gestor, Empresa: um.Empresa})
	}
	defer rows.Close()
	var obj PermissaoTipo
	for _, numero := range ps {
		obj = numero
	}
	return obj
}

//GetPermissaoTipoAll pegar todos os itens
func GetPermissaoTipoAll(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	permissaostipo, erro := getPermissaoTipoALL()

	if erro != nil {
		panic(erro)
	}

	json.NewEncoder(w).Encode(&permissaostipo)
}
func getPermissaoTipoALL() ([]PermissaoTipo, error) {

	rows, err := persistencia.DB.Query("SELECT id, nome, descricao, datacadastro, id_gestor, id_empresa FROM PermissaoTipo;")

	var ps []PermissaoTipo
	for rows.Next() {
		var um PermissaoTipo
		rows.Scan(&um.ID, &um.Nome, &um.Descricao, &um.DataCadastro, &um.Gestor.ID, &um.Empresa.ID)
		ps = append(ps, PermissaoTipo{ID: um.ID, Nome: um.Nome, Descricao: um.Descricao, DataCadastro: um.DataCadastro, Gestor: um.Gestor, Empresa: um.Empresa})
	}

	defer rows.Close()
	return ps, err
}
func primerioAcessoDadosPermissaoTipo(idGestor, idEmpresa int) {

	data, err := ioutil.ReadFile("controler/data/permissaotipo.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	erro := json.Unmarshal(data, &permissaostipo)
	fmt.Println("Teese", erro)
	for _, dados := range permissaostipo {

		var obj PermissaoTipo
		obj.DataCadastro = controler.PegarDataAtualString()
		obj.Nome = dados.Nome
		obj.Gestor.ID = idGestor
		obj.Empresa.ID = idEmpresa
		insertORUpPermissaoTipo(obj)
	}
}

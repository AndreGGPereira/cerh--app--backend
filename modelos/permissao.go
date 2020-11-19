package modelos

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andreggpereira/sisalicerce--app--backend/controler"
	"github.com/andreggpereira/sisalicerce--app--backend/persistencia"
	"github.com/gorilla/mux"
)

//Permissao tipos de permissão do Usuario
type Permissao struct {
	ID           int     `json:"id"`
	Nome         string  `json:"nome"`
	Descricao    string  `json:"descricao"`
	DataCadastro string  `json:"datacadastro"`
	Empresa      Empresa `json:"empresa"`
}

var permissoes []Permissao

//atualizar ou cadastrar
func insertORUpPermissao(obj Permissao) {

	//Se exister ID item deverá ser atualizado
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update permissao set  nome = $1, descricao = $2, datacadastro = $3, id_empresa = $4 where id =$5")
		stmt.Exec(obj.Nome, obj.Descricao, obj.DataCadastro, obj.Empresa.ID, obj.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into permissao(nome, descricao, datacadastro, id_empresa) values($1,$2,$3,$4)")
		stmt.Exec(obj.Nome, obj.Descricao, obj.DataCadastro, obj.Empresa.ID)

		if err != nil {
			log.Fatal("Cannot run insert statement", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	}
}

//CreatePermissao cadastrar permissao
func CreatePermissao(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Permissao
	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}
	obj.Empresa.ID, _ = controler.PegarIDEmpresaIDUsuario(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	insertORUpPermissao(obj)
	json.NewEncoder(w).Encode("Permissão cadastrado com sucesso")
}

//DeletePermissao deletar
func DeletePermissao(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Permissao

	obj.ID = controler.StringForInt(vars["id"])

	if obj.ID == 0 {
		fmt.Println("Não foi possível identificar a permissão")
	}

	error := deletarPermissao(obj.ID)
	if error != nil {
		json.NewEncoder(w).Encode(controler.Message{Message: "Operacao nao realizada", Status: 403})
	} else {
		json.NewEncoder(w).Encode(controler.Message{Message: "Permissao removida com sucesso", Status: 202})
	}

}

func deletarPermissao(id int) error {

	stmt, err := persistencia.DB.Prepare("delete from permissao where id = $1")
	stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	stmt.Close()
	return err
}

//GetPermissao busca item
func GetPermissao(w http.ResponseWriter, req *http.Request) {
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

func getPermissaoID(ID int) Permissao {

	rows, _ := persistencia.DB.Query("select id, nome, descricao, datacadastro from permissao where id = $1", ID)
	var ps []Permissao
	for rows.Next() {
		var um Permissao
		rows.Scan(&um.ID, &um.Nome, &um.Descricao, &um.DataCadastro)
		ps = append(ps, Permissao{ID: um.ID, Nome: um.Nome, Descricao: um.Descricao, DataCadastro: um.DataCadastro})
	}
	defer rows.Close()
	var obj Permissao
	for _, numero := range ps {
		obj = numero
	}
	return obj
}

//GetPermissaoAll pegar todos os itens
func GetPermissaoAll(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	permissoes, erro := getPermissaoALL()

	fmt.Println("lista permissoes", permissoes)
	if erro != nil {
		panic(erro)
	}
	json.NewEncoder(w).Encode(&permissoes)
}
func getPermissaoALL() ([]Permissao, error) {

	rows, _ := persistencia.DB.Query("SELECT id, nome, descricao, datacadastro FROM Permissao ORDER BY nome ASC")

	var ps []Permissao
	for rows.Next() {
		var um Permissao
		rows.Scan(&um.ID, &um.Nome, &um.Descricao, &um.DataCadastro)
		ps = append(ps, Permissao{ID: um.ID, Nome: um.Nome, Descricao: um.Descricao, DataCadastro: um.DataCadastro})
	}

	defer rows.Close()
	return ps, nil
}

// func getPermissaoALL() ([]Permissao, error) {

// 	rows, _ := persistencia.DB.Query("SELECT id, nome, descricao, datacadastro, id_empresa FROM Permissao")

// 	var ps []Permissao
// 	for rows.Next() {
// 		var um Permissao
// 		rows.Scan(&um.ID, &um.Nome, &um.Descricao, &um.DataCadastro)
// 		ps = append(ps, Permissao{ID: um.ID, Nome: um.Nome, Descricao: um.Descricao, DataCadastro: um.DataCadastro, Empresa: um.Empresa})
// 	}

// 	defer rows.Close()
// 	return ps, nil
// }

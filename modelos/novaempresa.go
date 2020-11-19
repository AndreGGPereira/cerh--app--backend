package modelos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type dados struct {
	Nome string `json:"nome"`
}

//var chan fdasa
//var msg chan string = make(chan string, 9)

func cadastroPadraoNovaEmpresa(idempresa, idusuario int) {

	go permissaoTipo(idempresa, idusuario)
	//	fmt.Println("Teste rotinas", <-msg)
}

func permissaoTipo(idempresa, idusuario int) {
	data, err := ioutil.ReadFile("/controle/data/permissao.json")
	if err != nil {
		fmt.Println(err)
	}
	var nomes []dados
	erro := json.Unmarshal(data, &nomes)
	fmt.Println("Nomes ", erro)
	for _, dados := range nomes {
		var obj PermissaoTipo
		obj.Empresa.ID = idempresa
		obj.Gestor.ID = idusuario
		obj.Nome = dados.Nome
		insertORUpPermissaoTipo(obj)

	}
	//	msg <- "permissaoTipo"
}

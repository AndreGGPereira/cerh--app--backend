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

//Avaliacao - CRUD
type Avaliacao struct {
	ID           int    `json:"id,omitempty"`
	DataCadastro string `json:"datacadastro,omitempty"`
	Login        string `json:"login,omitempty"`

	ValorTotalAvaliacoes float64 `json:"valortotalavaliacoes,omitempty"`

	FuncionarioAvaliado Funcionario `json:"funcionarioavaliado,omitempty"`

	FuncionarioAvaliador1  Funcionario `json:"funcionarioavaliador1,omitempty"`
	Avaliacaoefetuada1     bool        `json:"avaliacaoefetuada1,omitempty"`
	ValorTotalAvaliacao1   int         `json:"valortotalavaliacao1,omitempty"`
	ValorPesoAvaliacao1    int         `json:"valorpesoavaliacao1,omitempty"`
	DataAvaliacaoEfetuada1 string      `json:"dataavaliacaoefetuada1,omitempty"`

	FuncionarioAvaliador2  Funcionario `json:"funcionarioavaliador2,omitempty"`
	Avaliacaoefetuada2     bool        `json:"avaliacaoefetuada2,omitempty"`
	ValorTotalAvaliacao2   int         `json:"valortotalavaliacao2,omitempty"`
	ValorPesoAvaliacao2    int         `json:"valorpesoavaliacao2,omitempty"`
	DataAvaliacaoEfetuada2 string      `json:"dataavaliacaoefetuada2,omitempty"`

	FuncionarioAvaliador3  Funcionario `json:"funcionarioavaliador3,omitempty"`
	Avaliacaoefetuada3     bool        `json:"avaliacaoefetuada3,omitempty"`
	ValorTotalAvaliacao3   int         `json:"valortotalavaliacao3,omitempty"`
	ValorPesoAvaliacao3    int         `json:"valorpesoavaliacao3,omitempty"`
	DataAvaliacaoEfetuada3 string      `json:"dataavaliacaoefetuada3,omitempty"`

	FuncionarioAvaliador4  Funcionario `json:"funcionarioavaliador4,omitempty"`
	Avaliacaoefetuada4     bool        `json:"avaliacaoefetuada4,omitempty"`
	ValorTotalAvaliacao4   int         `json:"valortotalavaliacao4,omitempty"`
	ValorPesoAvaliacao4    int         `json:"valorpesoavaliacao4,omitempty"`
	DataAvaliacaoEfetuada4 string      `json:"dataavaliacaoefetuada4,omitempty"`

	Gestor  Gestor  `json:"gestor,omitempty"`
	Empresa Empresa `json:"empresa,omitempty"`
	Message controler.Message
}

//AvaliacaoList struct consulta
type AvaliacaoList struct {
	Avaliacao []Avaliacao       `json:"avaliacao"`
	Contador  int               `json:"contador"`
	Message   controler.Message `json:"message,omitempty"`
}

var avaliacoes []Avaliacao

func insertORUpAvaliacao(obj Avaliacao) error {

	var err error
	if obj.ID != 0 {

		stmt, err := persistencia.DB.Prepare("update cargo set nome = $1, datacadastro = $2, id_gestor = $3, id_empresa = $4 where id =$5")
		stmt.Exec(obj.ID, obj.DataCadastro, obj.Login, obj.ValorTotalAvaliacoes, obj.FuncionarioAvaliado.ID, obj.FuncionarioAvaliador1.ID, obj.Avaliacaoefetuada1, obj.ValorTotalAvaliacao1, obj.ValorPesoAvaliacao1, obj.DataAvaliacaoEfetuada1, obj.FuncionarioAvaliador2.ID, obj.Avaliacaoefetuada2, obj.ValorTotalAvaliacao2, obj.ValorPesoAvaliacao2, obj.DataAvaliacaoEfetuada2, obj.FuncionarioAvaliador3.ID, obj.Avaliacaoefetuada3, obj.ValorTotalAvaliacao3, obj.ValorPesoAvaliacao3, obj.DataAvaliacaoEfetuada3, obj.FuncionarioAvaliador4.ID, obj.Avaliacaoefetuada4, obj.ValorTotalAvaliacao4, obj.ValorPesoAvaliacao4, obj.DataAvaliacaoEfetuada4, obj.Gestor.ID, obj.Empresa.ID)

		if err != nil {
			log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()

	} else {

		stmt, err := persistencia.DB.Prepare("insert into avaliacao(datacadastro, login, valortotalavaliacoes, id_funcionarioavaliado, id_funcionarioavaliador1, avaliacaoefetuada1, valortotalavaliacao1, valorpesoavaliacao1, dataavaliacaoefetuada1, id_funcionarioavaliador2, avaliacaoefetuada2, valortotalavaliacao2, valorpesoavaliacao2, dataavaliacaoefetuada2, id_funcionarioavaliador3, avaliacaoefetuada3, valortotalavaliacao3, valorpesoavaliacao3, dataavaliacaoefetuada3, id_funcionarioavaliador4, avaliacaoefetuada4, valortotalavaliacao4, valorpesoavaliacao4, dataavaliacaoefetuada4, id_gestor, id_empresa) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26) RETURNING id")
		stmt.Exec(obj.DataCadastro, obj.Login, obj.ValorTotalAvaliacoes, obj.FuncionarioAvaliado.ID, obj.FuncionarioAvaliador1.ID, obj.Avaliacaoefetuada1, obj.ValorTotalAvaliacao1, obj.ValorPesoAvaliacao1, obj.DataAvaliacaoEfetuada1, obj.FuncionarioAvaliador2.ID, obj.Avaliacaoefetuada2, obj.ValorTotalAvaliacao2, obj.ValorPesoAvaliacao2, obj.DataAvaliacaoEfetuada2, obj.FuncionarioAvaliador3.ID, obj.Avaliacaoefetuada3, obj.ValorTotalAvaliacao3, obj.ValorPesoAvaliacao3, obj.DataAvaliacaoEfetuada3, obj.FuncionarioAvaliador4.ID, obj.Avaliacaoefetuada4, obj.ValorTotalAvaliacao4, obj.ValorPesoAvaliacao4, obj.DataAvaliacaoEfetuada4, obj.Gestor.ID, obj.Empresa.ID)

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

//CreateAvaliacao contas
func CreateAvaliacao(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)

	var obj Avaliacao

	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	obj.DataCadastro = controler.PegarDataAtualStringNew()

	if obj.ID == 0 {
		obj.ValorTotalAvaliacoes = 0.0
		obj.Avaliacaoefetuada1 = false
		obj.DataAvaliacaoEfetuada1 = "1970-01-01 00:00:01"
		obj.DataAvaliacaoEfetuada2 = "1970-01-01 00:00:01"
		obj.DataAvaliacaoEfetuada3 = "1970-01-01 00:00:01"
		obj.DataAvaliacaoEfetuada4 = "1970-01-01 00:00:01"
		obj.FuncionarioAvaliado = obj.FuncionarioAvaliador1
	}

	fmt.Println("obj.FuncionarioAvaliado.ID", obj.FuncionarioAvaliado.ID)
	fmt.Println(" obj.FuncionarioAvaliado1.ID", obj.FuncionarioAvaliador1.ID)
	fmt.Println("obj.FuncionarioAvaliado2.ID", obj.FuncionarioAvaliador2.ID)
	fmt.Println("obj.FuncionarioAvaliado3.ID", obj.FuncionarioAvaliador3.ID)
	fmt.Println("obj.FuncionarioAvaliado4.ID", obj.FuncionarioAvaliador4.ID)
	fmt.Println(" obj.Gestor.ID", obj.Gestor.ID)
	fmt.Println("obj.Empresa.ID", obj.Empresa.ID)

	fmt.Println(" valorpesoavaliacao1", obj.ValorPesoAvaliacao1)
	fmt.Println("valorpesoavaliacao2", obj.ValorPesoAvaliacao2)
	fmt.Println(" valorpesoavaliacao3", obj.ValorPesoAvaliacao3)
	fmt.Println("valorpesoavaliacao4", obj.ValorPesoAvaliacao4)

	err = insertORUpAvaliacao(obj)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Cadastro realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}

//DeleteAvaliacao Avaliacao
func DeleteAvaliacao(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Avaliacao

	obj.ID = controler.StringForInt(vars["id"])

	if obj.ID == 0 {
		fmt.Println("Não foi possível encontrar a Etapa da Obra")
	}

	err := deletarAvaliacao(obj.ID)

	if err != nil {
		obj.Message.Status = 202
		obj.Message.Message = "Remocao realizado com sucesso!!"

	} else {
		obj.Message.Status = 304
		obj.Message.Message = "Nao foi possivel relizar a operacao"
	}

	json.NewEncoder(w).Encode(obj.Message)
}
func deletarAvaliacao(id int) error {
	stmt, err := persistencia.DB.Prepare("delete from avalicao where id = $1")
	stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	return err
}

//GetAvaliacao Cargo
func GetAvaliacao(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var obj Cargo
	var err error
	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Nao foi possivel encontrar o item")
	}
	obj, err = getCargoID(obj.ID)

	json.NewEncoder(w).Encode(&obj)
}

func getAvaliacaoID(ID int) (Avaliacao, error) {

	rows, err := persistencia.DB.Query("SELECT id, datacadastro, login, valortotalavaliacoes, id_funcionarioavaliado, id_funcionarioavaliador1, avaliacaoefetuada1, valortotalavaliacao1, valorpesoavaliacao1, dataavaliacaoefetuada1, id_funcionarioavaliador2, avaliacaoefetuada2, valortotalavaliacao2, valorpesoavaliacao2, dataavaliacaoefetuada2, id_funcionarioavaliador3, avaliacaoefetuada3, valortotalavaliacao3, valorpesoavaliacao3, dataavaliacaoefetuada3, id_funcionarioavaliador4, avaliacaoefetuada4, valortotalavaliacao4, valorpesoavaliacao4, dataavaliacaoefetuada4, id_gestor, id_empresa FROM avaliacao where id = $1 ORDER BY datacadastro DESC ", ID)
	if err != nil {
		log.Fatal("Erro getCargoID", err)
	}
	var ps []Avaliacao
	for rows.Next() {
		var obj Avaliacao
		rows.Scan(&obj.ID, &obj.DataCadastro, &obj.Login, &obj.ValorTotalAvaliacoes, &obj.FuncionarioAvaliado.ID, &obj.FuncionarioAvaliador1.ID, &obj.Avaliacaoefetuada1, &obj.ValorTotalAvaliacao1, &obj.ValorPesoAvaliacao1, &obj.DataAvaliacaoEfetuada1, &obj.FuncionarioAvaliador2.ID, &obj.Avaliacaoefetuada2, &obj.ValorTotalAvaliacao2, &obj.ValorPesoAvaliacao2, &obj.DataAvaliacaoEfetuada2, &obj.FuncionarioAvaliador3.ID, &obj.Avaliacaoefetuada3, &obj.ValorTotalAvaliacao3, &obj.ValorPesoAvaliacao3, &obj.DataAvaliacaoEfetuada3, &obj.FuncionarioAvaliador4.ID, &obj.Avaliacaoefetuada4, &obj.ValorTotalAvaliacao4, &obj.ValorPesoAvaliacao4, &obj.DataAvaliacaoEfetuada4, &obj.Gestor.ID, &obj.Empresa.ID)
		ps = append(ps, Avaliacao{ID: obj.ID, DataCadastro: obj.DataCadastro, Login: obj.Login, ValorTotalAvaliacoes: obj.ValorTotalAvaliacoes, FuncionarioAvaliado: obj.FuncionarioAvaliado, FuncionarioAvaliador1: obj.FuncionarioAvaliador1, Avaliacaoefetuada1: obj.Avaliacaoefetuada1, ValorTotalAvaliacao1: obj.ValorTotalAvaliacao1, ValorPesoAvaliacao1: obj.ValorPesoAvaliacao1, DataAvaliacaoEfetuada1: obj.DataAvaliacaoEfetuada1, FuncionarioAvaliador2: obj.FuncionarioAvaliador2, Avaliacaoefetuada2: obj.Avaliacaoefetuada2, ValorTotalAvaliacao2: obj.ValorTotalAvaliacao2, ValorPesoAvaliacao2: obj.ValorPesoAvaliacao2, DataAvaliacaoEfetuada2: obj.DataAvaliacaoEfetuada2, FuncionarioAvaliador3: obj.FuncionarioAvaliador3, Avaliacaoefetuada3: obj.Avaliacaoefetuada3, ValorTotalAvaliacao3: obj.ValorTotalAvaliacao3, ValorPesoAvaliacao3: obj.ValorPesoAvaliacao3, DataAvaliacaoEfetuada3: obj.DataAvaliacaoEfetuada3, FuncionarioAvaliador4: obj.FuncionarioAvaliador4, Avaliacaoefetuada4: obj.Avaliacaoefetuada4, ValorTotalAvaliacao4: obj.ValorTotalAvaliacao4, ValorPesoAvaliacao4: obj.ValorPesoAvaliacao4, DataAvaliacaoEfetuada4: obj.DataAvaliacaoEfetuada4, Gestor: obj.Gestor, Empresa: obj.Empresa})
	}
	defer rows.Close()
	var obj Avaliacao
	for _, nobjero := range ps {
		obj = nobjero
	}
	return obj, err
}

//GetAvaliacaoAll Cargo todos
func GetAvaliacaoAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var obj Cargo
	var objList CargoList
	var erro error
	obj.Empresa.ID, obj.Gestor.ID = controler.PegarIDEmpresaIDGestor(req)
	objList.Cargo, erro = getCargoAll(obj.Empresa.ID)
	objList.Contador = len(objList.Cargo)

	if erro != nil {
		objList.Message.Message = " Nao a itens na lista"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 304
	}

	json.NewEncoder(w).Encode(&objList)

}
func getAvaliacaoAll(id int) ([]Avaliacao, error) {

	rows, err := persistencia.DB.Query("SELECT id, datacadastro, login, valortotalavaliacoes, id_funcionarioavaliado, id_funcionarioavaliador1, avaliacaoefetuada1, valortotalavaliacao1, valorpesoavaliacao1, dataavaliacaoefetuada1, id_funcionarioavaliador2, avaliacaoefetuada2, valortotalavaliacao2, valorpesoavaliacao2, dataavaliacaoefetuada2, id_funcionarioavaliador3, avaliacaoefetuada3, valortotalavaliacao3, valorpesoavaliacao3, dataavaliacaoefetuada3, id_funcionarioavaliador4, funcionarioavaliador4, avaliacaoefetuada4, valortotalavaliacao4, valorpesoavaliacao4, dataavaliacaoefetuada4, id_gestor, id_empresa FROM avaliacao where id_empresa = $1 ORDER BY datacadastro DESC ", id)

	if err != nil {
		log.Fatal("Erro getAvaliacaoAll", err)
	}
	var ps []Avaliacao
	for rows.Next() {
		var obj Avaliacao
		rows.Scan(&obj.ID, &obj.DataCadastro, &obj.Login, &obj.ValorTotalAvaliacoes, &obj.FuncionarioAvaliado.ID, &obj.FuncionarioAvaliador1.ID, &obj.Avaliacaoefetuada1, &obj.ValorTotalAvaliacao1, &obj.ValorPesoAvaliacao1, &obj.DataAvaliacaoEfetuada1, &obj.FuncionarioAvaliador2.ID, &obj.Avaliacaoefetuada2, &obj.ValorTotalAvaliacao2, &obj.ValorPesoAvaliacao2, &obj.DataAvaliacaoEfetuada2, &obj.FuncionarioAvaliador3.ID, &obj.Avaliacaoefetuada3, &obj.ValorTotalAvaliacao3, &obj.ValorPesoAvaliacao3, &obj.DataAvaliacaoEfetuada3, &obj.FuncionarioAvaliador4.ID, &obj.Avaliacaoefetuada4, &obj.ValorTotalAvaliacao4, &obj.ValorPesoAvaliacao4, &obj.DataAvaliacaoEfetuada4, &obj.Gestor.ID, &obj.Empresa.ID)
		ps = append(ps, Avaliacao{ID: obj.ID, DataCadastro: obj.DataCadastro, Login: obj.Login, ValorTotalAvaliacoes: obj.ValorTotalAvaliacoes, FuncionarioAvaliado: obj.FuncionarioAvaliado, FuncionarioAvaliador1: obj.FuncionarioAvaliador1, Avaliacaoefetuada1: obj.Avaliacaoefetuada1, ValorTotalAvaliacao1: obj.ValorTotalAvaliacao1, ValorPesoAvaliacao1: obj.ValorPesoAvaliacao1, DataAvaliacaoEfetuada1: obj.DataAvaliacaoEfetuada1, FuncionarioAvaliador2: obj.FuncionarioAvaliador2, Avaliacaoefetuada2: obj.Avaliacaoefetuada2, ValorTotalAvaliacao2: obj.ValorTotalAvaliacao2, ValorPesoAvaliacao2: obj.ValorPesoAvaliacao2, DataAvaliacaoEfetuada2: obj.DataAvaliacaoEfetuada2, FuncionarioAvaliador3: obj.FuncionarioAvaliador3, Avaliacaoefetuada3: obj.Avaliacaoefetuada3, ValorTotalAvaliacao3: obj.ValorTotalAvaliacao3, ValorPesoAvaliacao3: obj.ValorPesoAvaliacao3, DataAvaliacaoEfetuada3: obj.DataAvaliacaoEfetuada3, FuncionarioAvaliador4: obj.FuncionarioAvaliador4, Avaliacaoefetuada4: obj.Avaliacaoefetuada4, ValorTotalAvaliacao4: obj.ValorTotalAvaliacao4, ValorPesoAvaliacao4: obj.ValorPesoAvaliacao4, DataAvaliacaoEfetuada4: obj.DataAvaliacaoEfetuada4, Gestor: obj.Gestor, Empresa: obj.Empresa})
	}
	defer rows.Close()
	return ps, err
}

//GetAvaliacaoPGAll pegar todos os itens
func GetAvaliacaoPGAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	var page int
	var obj Avaliacao
	var objList AvaliacaoList
	var err error

	obj.Empresa.ID, _ = controler.PegarIDEmpresaIDGestor(req)
	page = controler.StringForInt(vars["page"])

	objList.Avaliacao, err = getAvaliacaoPGAll(obj.Empresa.ID, page)
	objList.Contador = getCountAvaliacaoAll(obj.Empresa.ID)

	if err != nil {
		objList.Message.Message = "Nao possivel realizar a consulta"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 200

	}

	json.NewEncoder(w).Encode(&objList)
}

func getAvaliacaoPGAll(idempresa, page int) ([]Avaliacao, error) {
	rows, err := persistencia.DB.Query("SELECT a.id, a.datacadastro, a.login, a.valortotalavaliacoes, a.id_funcionarioavaliado, a.id_funcionarioavaliador1, a.avaliacaoefetuada1, a.valortotalavaliacao1, a.valorpesoavaliacao1, a.dataavaliacaoefetuada1, a.id_funcionarioavaliador2, a.avaliacaoefetuada2, a.valortotalavaliacao2, a.valorpesoavaliacao2, a.dataavaliacaoefetuada2, a.id_funcionarioavaliador3, a.avaliacaoefetuada3, a.valortotalavaliacao3, a.valorpesoavaliacao3, a.dataavaliacaoefetuada3, a.id_funcionarioavaliador4, a.avaliacaoefetuada4, a.valortotalavaliacao4, a.valorpesoavaliacao4, a.dataavaliacaoefetuada4, a.id_gestor, a.id_empresa, fu.nome FROM avaliacao AS a INNER JOIN funcionario as fu ON a.id_funcionarioavaliado = fu.id WHERE a.id_empresa =$1 LIMIT $2 OFFSET $3 ", idempresa, 10, page)

	if err != nil {
		fmt.Println(err)
	}

	var ps []Avaliacao
	for rows.Next() {
		var obj Avaliacao
		rows.Scan(&obj.ID, &obj.DataCadastro, &obj.Login, &obj.ValorTotalAvaliacoes, &obj.FuncionarioAvaliado.ID, &obj.FuncionarioAvaliador1.ID, &obj.Avaliacaoefetuada1, &obj.ValorTotalAvaliacao1, &obj.ValorPesoAvaliacao1, &obj.DataAvaliacaoEfetuada1, &obj.FuncionarioAvaliador2.ID, &obj.Avaliacaoefetuada2, &obj.ValorTotalAvaliacao2, &obj.ValorPesoAvaliacao2, &obj.DataAvaliacaoEfetuada2, &obj.FuncionarioAvaliador3.ID, &obj.Avaliacaoefetuada3, &obj.ValorTotalAvaliacao3, &obj.ValorPesoAvaliacao3, &obj.DataAvaliacaoEfetuada3, &obj.FuncionarioAvaliador4.ID, &obj.Avaliacaoefetuada4, &obj.ValorTotalAvaliacao4, &obj.ValorPesoAvaliacao4, &obj.DataAvaliacaoEfetuada4, &obj.Gestor.ID, &obj.Empresa.ID, &obj.FuncionarioAvaliado.Nome)
		ps = append(ps, Avaliacao{ID: obj.ID, DataCadastro: obj.DataCadastro, Login: obj.Login, ValorTotalAvaliacoes: obj.ValorTotalAvaliacoes, FuncionarioAvaliado: obj.FuncionarioAvaliado, FuncionarioAvaliador1: obj.FuncionarioAvaliador1, Avaliacaoefetuada1: obj.Avaliacaoefetuada1, ValorTotalAvaliacao1: obj.ValorTotalAvaliacao1, ValorPesoAvaliacao1: obj.ValorPesoAvaliacao1, DataAvaliacaoEfetuada1: obj.DataAvaliacaoEfetuada1, FuncionarioAvaliador2: obj.FuncionarioAvaliador2, Avaliacaoefetuada2: obj.Avaliacaoefetuada2, ValorTotalAvaliacao2: obj.ValorTotalAvaliacao2, ValorPesoAvaliacao2: obj.ValorPesoAvaliacao2, DataAvaliacaoEfetuada2: obj.DataAvaliacaoEfetuada2, FuncionarioAvaliador3: obj.FuncionarioAvaliador3, Avaliacaoefetuada3: obj.Avaliacaoefetuada3, ValorTotalAvaliacao3: obj.ValorTotalAvaliacao3, ValorPesoAvaliacao3: obj.ValorPesoAvaliacao3, DataAvaliacaoEfetuada3: obj.DataAvaliacaoEfetuada3, FuncionarioAvaliador4: obj.FuncionarioAvaliador4, Avaliacaoefetuada4: obj.Avaliacaoefetuada4, ValorTotalAvaliacao4: obj.ValorTotalAvaliacao4, ValorPesoAvaliacao4: obj.ValorPesoAvaliacao4, DataAvaliacaoEfetuada4: obj.DataAvaliacaoEfetuada4, Gestor: obj.Gestor, Empresa: obj.Empresa})
	}

	defer rows.Close()
	return ps, err
}

func getCountAvaliacaoAll(idempresa int) int {

	rows, err := persistencia.DB.Query("SELECT COUNT(id) FROM avaliacao where id_empresa = $1 ", idempresa)

	if err != nil {
		fmt.Println(err)
	}
	var ps []Avaliacao
	for rows.Next() {
		var obj Avaliacao
		rows.Scan(&obj.ID)
		ps = append(ps, Avaliacao{ID: obj.ID})
	}

	defer rows.Close()

	var contador int
	for _, nobjero := range ps {
		contador = nobjero.ID
	}

	fmt.Println(" Dados do contator ", contador)
	return contador
}

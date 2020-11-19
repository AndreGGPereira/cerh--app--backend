package modelos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andreggpereira/cerh--app--backend/controler"
	"github.com/andreggpereira/cerh--app--backend/persistencia"
	"github.com/gorilla/mux"
)

//QuestionarioAvaliacao capitalize to export from package
type QuestionarioAvaliacao struct {
	ID                 int    `json:"id,omitempty"`
	Nome               string `json:"nome,omitempty"`
	DataCadastro       string `json:"datacadastro,omitempty"`
	DataQuestionario   string `json:"dataquestionario,omitempty"`
	DataResposta       string `json:"dataresposta,omitempty"`
	Auditoria          string `json:"auditoria,omitempty"`
	Respondiada        bool   `json:"respondida,omitempty"`
	PontuacaoFinal     int    `json:"pontuacaofinal,omitempty"`
	Peso               int    `json:"peso,omitempty"`
	ValorPesoAvaliacao int    `json:"valorpesoavaliacao,omitempty"`

	Avaliacao            Avaliacao   `json:"avaliacao,omitempty"`
	FuncionarioAvaliado  Funcionario `json:"funcionarioavaliado,omitempty"`
	FuncionarioAvaliador Funcionario `json:"funcionarioavaliador,omitempty"`
	Empresa              Empresa     `json:"empresa,omitempty"`
}

var questionarioAvaliacaoes []QuestionarioAvaliacao

//QuestionarioAvaliacaoList struct consulta
type QuestionarioAvaliacaoList struct {
	QuestionarioAvaliacao []QuestionarioAvaliacao `json:"questionarioavaliacao,omitempty"`
	Contador              int                     `json:"contador,omitempty"`
	Message               controler.Message       `json:"message,omitempty"`
}

func insertORUpQuestionarioAvaliacao(obj QuestionarioAvaliacao) int {

	var id int
	//Se exister ID item deverá ser atualizado
	if obj.ID != 0 {
		id = obj.ID
		stmt, err := persistencia.DB.Prepare("update QuestionarioAvaliacao set nome = $1, datacadastro = $2, dataquestionario = $3, dataresposta = $4, auditoria = $5, respondida = $6, pontuacaofinal = $7, peso = $8, valorpesoavaliacao = $9, id_avaliacao = $10, id_funcionarioavaliado = $11, id_funcionarioavaliador = $12, id_empresa = $13 where id =$14")
		err = stmt.QueryRow(obj.Nome, obj.DataCadastro, obj.DataQuestionario, obj.DataResposta, obj.Auditoria, obj.Respondiada, obj.PontuacaoFinal, obj.Peso, obj.ValorPesoAvaliacao, obj.Avaliacao.ID, obj.FuncionarioAvaliado.ID, obj.FuncionarioAvaliador.ID, obj.Empresa.ID, obj.ID).Scan(&id)
		if err != nil {
			//log.Fatal("Não foi possível atualizar o item", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	} else {

		stmt, err := persistencia.DB.Prepare("insert into QuestionarioAvaliacao(nome, datacadastro, dataquestionario, dataresposta, auditoria, respondida, pontuacaofinal, peso, valorpesoavaliacao, id_avaliacao, id_funcionarioavaliado, id_funcionarioavaliador, id_empresa) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)  RETURNING id")
		err = stmt.QueryRow(obj.Nome, obj.DataCadastro, obj.DataQuestionario, obj.DataResposta, obj.Auditoria, obj.Respondiada, obj.PontuacaoFinal, obj.Peso, obj.ValorPesoAvaliacao, obj.Avaliacao.ID, obj.FuncionarioAvaliado.ID, obj.FuncionarioAvaliador.ID, obj.Empresa.ID).Scan(&id)
		fmt.Println("Teste ID ", &id, id)
		if err != nil {
			//log.Fatal("Cannot run insert statement", err)
			fmt.Println(err)
		} else {
			fmt.Println("Operação concluida com sucesso!!")
		}
		defer stmt.Close()
	}
	return id
}

//CreateQuestionarioAvaliacao cadastro
func CreateQuestionarioAvaliacao(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var obj QuestionarioAvaliacao
	err := decoder.Decode(&obj)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	}

	obj.DataCadastro = controler.PegarDataAtualStringNew()

	obj.Empresa.ID, _ = controler.PegarIDEmpresaIDUsuario(req)

	insertORUpQuestionarioAvaliacao(obj)
	json.NewEncoder(w).Encode("QuestionarioAvaliacao cadastrado com sucesso")
}

//GetQuestionarioAvaliacaoAll pegar todos os itens
func GetQuestionarioAvaliacaoAll(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var obj QuestionarioAvaliacao
	var objList QuestionarioAvaliacaoList
	var err error
	obj.Empresa.ID, _ = controler.PegarIDEmpresaIDUsuario(req)

	objList.QuestionarioAvaliacao, err = getQuestionarioAvaliacaoALL(obj.Empresa.ID)
	objList.Contador = len(objList.QuestionarioAvaliacao)

	if err != nil {
		objList.Message.Message = "Nao possivel realizar a consulta"
		objList.Message.Status = 304
	} else {
		objList.Message.Message = "Consulta realizada com sucesso"
		objList.Message.Status = 200
	}

	json.NewEncoder(w).Encode(&objList)
}
func getQuestionarioAvaliacaoALL(idempresa int) ([]QuestionarioAvaliacao, error) {
	rows, err := persistencia.DB.Query("SELECT id, nome, datacadastro, dataquestionario, dataresposta, auditoria, respondida, pontuacaofinal, peso, valorpesoavaliacao, id_avaliacao, id_funcionarioavaliado, id_funcionarioavaliador, id_empresa FROM QuestionarioAvaliacao WHERE id_empresa =$1", idempresa)

	if err != nil {
		fmt.Println(err)
	}

	var ps []QuestionarioAvaliacao
	for rows.Next() {
		var obj QuestionarioAvaliacao
		rows.Scan(&obj.ID, &obj.Nome, &obj.DataCadastro, &obj.DataQuestionario, &obj.DataResposta, &obj.Auditoria, &obj.Respondiada, &obj.PontuacaoFinal, &obj.Peso, &obj.ValorPesoAvaliacao, &obj.Avaliacao.ID, &obj.FuncionarioAvaliado.ID, &obj.FuncionarioAvaliador.ID, &obj.Empresa.ID)
		ps = append(ps, QuestionarioAvaliacao{ID: obj.ID, Nome: obj.Nome, DataCadastro: obj.DataCadastro, DataQuestionario: obj.DataQuestionario, DataResposta: obj.DataResposta, Auditoria: obj.Auditoria, Respondiada: obj.Respondiada, PontuacaoFinal: obj.PontuacaoFinal, Peso: obj.Peso, ValorPesoAvaliacao: obj.ValorPesoAvaliacao, Avaliacao: obj.Avaliacao, FuncionarioAvaliado: obj.FuncionarioAvaliado, FuncionarioAvaliador: obj.FuncionarioAvaliador, Empresa: obj.Empresa})
	}

	defer rows.Close()
	return ps, err
}

//DeleteQuestionarioAvaliacao deleteFornecedor
func DeleteQuestionarioAvaliacao(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var obj QuestionarioAvaliacao
	var err error
	vars := mux.Vars(req)

	obj.ID, err = strconv.Atoi(vars["id"])

	if err != nil {
		json.NewEncoder(w).Encode("Erro ao inserir o identificador")
	}

	deletarQuestionarioAvaliacao(obj.ID)

	//json.NewEncoder(w).Encode(QuestionarioAvaliacaos)
}

//DeletarQuestionarioAvaliacao deletar usuaria
func deletarQuestionarioAvaliacao(id int) {

	stmt2, err := persistencia.DB.Prepare("delete from questionarioavaliacao where id = $1")
	stmt2.Exec(id)

	if err != nil {
		panic(err)
	}
	stmt2.Close()
}

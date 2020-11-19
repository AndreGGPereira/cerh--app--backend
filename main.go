package main

import (
	//	"crypto/hmac"
	//	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"strings"

	//	"html/template"
	_ "github.com/lib/pq"

	"log"
	"net/http"
	"net/smtp"
	"strconv"

	//	"strings"
	//	"time"
	//	"encoding/json"HEROKU
	//"html/template"
	"os"
	//"strconv"
	//"log"
	"github.com/andreggpereira/cerh--app--backend/controler"
	"github.com/andreggpereira/cerh--app--backend/modelos"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

//variavel com ponteiro do template
func main() {

	fmt.Println("Hello, playground")
	router := mux.NewRouter()
	c := cors.AllowAll()

	//criarBanco()
	//sendEmailTeste()
	// modelos.AddProdutoOF()
	router.HandleFunc("/", index)
	//router.HandleFunc("/login", modelos.UsuarioLogin).Methods("POST")
	router.HandleFunc("/login", modelos.GestorLogin).Methods("POST")
	// router.HandleFunc("/empresa", modelos.CreateEmpresa).Methods("POST")
	router.HandleFunc("/empresa", modelos.CreateEmpresa).Methods("POST")
	router.HandleFunc("/buscacep/{cep}", modelos.BuscaCep).Methods("GET")
	router.HandleFunc("/buscaremail", modelos.BuscarEmail).Methods("POST")
	router.HandleFunc("/validaremail/", modelos.ValidarEmailToken).Methods("GET")

	// router.HandleFunc("/cliente/", ValidateMiddleware(modelos.GetClienteAll)).Methods("GET")
	// router.HandleFunc("/cliente/{id}", ValidateMiddleware(modelos.GetClientePGAll)).Methods("GET")
	// router.HandleFunc("/clienteid/{id}", ValidateMiddleware(modelos.GetCliente)).Methods("GET")
	// router.HandleFunc("/cliente/", ValidateMiddleware(modelos.CreateCliente)).Methods("POST")
	// router.HandleFunc("/cliente/{id}", ValidateMiddleware(modelos.DeleteCliente)).Methods("DELETE")
	// router.HandleFunc("/cliente/{id}", ValidateMiddleware(modelos.CreateCliente)).Methods("PUT")
	// router.HandleFunc("/clienteSele", ValidateMiddleware(modelos.GetClienteConsulta)).Methods("PUT")

	// router.HandleFunc("/movimentacaotipo/", ValidateMiddleware(modelos.GetMovimentacaoTipoAll)).Methods("GET")
	// router.HandleFunc("/movimentacaotipo/{id}", ValidateMiddleware(modelos.GetMovimentacaoTipo)).Methods("GET")
	// router.HandleFunc("/movimentacaotipo/", ValidateMiddleware(modelos.CreateMovimentacaoTipo)).Methods("POST")
	// router.HandleFunc("/movimentacaotipo/{id}", ValidateMiddleware(modelos.DeleteMovimentacaoTipo)).Methods("DELETE")

	// router.HandleFunc("/notificacoes/", ValidateMiddleware(modelos.GetMovimentacaoTipoAll)).Methods("GET")
	// router.HandleFunc("/notificacoes/{id}", ValidateMiddleware(modelos.GetMovimentacaoTipo)).Methods("GET")
	// router.HandleFunc("/notificacoes/", ValidateMiddleware(modelos.CreateMovimentacaoTipo)).Methods("POST")
	// router.HandleFunc("/notificacoes/{id}", ValidateMiddleware(modelos.DeleteMovimentacaoTipo)).Methods("DELETE")

	// router.HandleFunc("/of/", ValidateMiddleware(modelos.GetOFAll)).Methods("GET")
	// router.HandleFunc("/of/{page}", ValidateMiddleware(modelos.GetOFAll)).Methods("GET")
	// router.HandleFunc("/ofconsulta/{cliente}/{obra}/{datainicio}/{datafim}/{pgconcluido}/{concluido}", ValidateMiddleware(modelos.GetOFConsulta)).Methods("GET")
	// router.HandleFunc("/ofconsultadate/{datainicio}/{datafim}", ValidateMiddleware(modelos.GetOFConsultaDate)).Methods("GET")

	router.HandleFunc("/avaliacaopg/{page}", ValidateMiddleware(modelos.GetAvaliacaoPGAll)).Methods("GET")
	router.HandleFunc("/avalicao/", ValidateMiddleware(modelos.GetAvaliacaoAll)).Methods("GET")
	router.HandleFunc("/avaliacao/{id}", ValidateMiddleware(modelos.GetAvaliacao)).Methods("GET")
	router.HandleFunc("/avaliacao/", ValidateMiddleware(modelos.CreateAvaliacao)).Methods("POST")
	router.HandleFunc("/avaliacao/{id}", ValidateMiddleware(modelos.DeleteAvaliacao)).Methods("DELETE")
	router.HandleFunc("/avaliacao/{id}", ValidateMiddleware(modelos.CreateAvaliacao)).Methods("PUT")

	router.HandleFunc("/cargo/", ValidateMiddleware(modelos.GetCargoAll)).Methods("GET")
	router.HandleFunc("/cargo/{id}", ValidateMiddleware(modelos.GetCargo)).Methods("GET")
	router.HandleFunc("/cargo/", ValidateMiddleware(modelos.CreateCargo)).Methods("POST")
	router.HandleFunc("/cargo/{id}", ValidateMiddleware(modelos.DeleteCargo)).Methods("DELETE")
	router.HandleFunc("/cargo/{id}", ValidateMiddleware(modelos.CreateCargo)).Methods("PUT")

	router.HandleFunc("/departamento/", ValidateMiddleware(modelos.GetDepartamentoAll)).Methods("GET")
	router.HandleFunc("/departamento/{id}", ValidateMiddleware(modelos.GetDepartamento)).Methods("GET")
	router.HandleFunc("/departamento/", ValidateMiddleware(modelos.CreateDepartamento)).Methods("POST")
	router.HandleFunc("/departamento/{id}", ValidateMiddleware(modelos.DeleteDepartamento)).Methods("DELETE")
	router.HandleFunc("/departamento/{id}", ValidateMiddleware(modelos.CreateDepartamento)).Methods("PUT")

	router.HandleFunc("/funcao/", ValidateMiddleware(modelos.GetFuncaoAll)).Methods("GET")
	router.HandleFunc("/funcao/{id}", ValidateMiddleware(modelos.GetFuncao)).Methods("GET")
	router.HandleFunc("/funcao/", ValidateMiddleware(modelos.CreateFuncao)).Methods("POST")
	router.HandleFunc("/funcao/{id}", ValidateMiddleware(modelos.DeleteFuncao)).Methods("DELETE")
	router.HandleFunc("/funcao/{id}", ValidateMiddleware(modelos.CreateFuncao)).Methods("PUT")

	router.HandleFunc("/grauinstrucao/", ValidateMiddleware(modelos.GetGrauinstrucaoAll)).Methods("GET")
	router.HandleFunc("/grauinstrucao/{id}", ValidateMiddleware(modelos.GetGrauinstrucao)).Methods("GET")
	router.HandleFunc("/grauinstrucao/", ValidateMiddleware(modelos.CreateGrauinstrucao)).Methods("POST")
	router.HandleFunc("/grauinstrucao/{id}", ValidateMiddleware(modelos.DeleteGrauinstrucao)).Methods("DELETE")
	router.HandleFunc("/grauinstrucao/{id}", ValidateMiddleware(modelos.CreateGrauinstrucao)).Methods("PUT")

	router.HandleFunc("/permissao/", ValidateMiddleware(modelos.GetPermissaoAll)).Methods("GET")
	router.HandleFunc("/permissao/{id}", ValidateMiddleware(modelos.GetPermissao)).Methods("GET")
	router.HandleFunc("/permissao/", ValidateMiddleware(modelos.CreatePermissao)).Methods("POST")
	router.HandleFunc("/permissao/{id}", ValidateMiddleware(modelos.DeletePermissao)).Methods("DELETE")
	router.HandleFunc("/permissao/{id}", ValidateMiddleware(modelos.CreatePermissao)).Methods("PUT")

	router.HandleFunc("/competenciaforjob/{id}", ValidateMiddleware(modelos.GetCompetenciaForJob)).Methods("GET")
	router.HandleFunc("/competencia/", ValidateMiddleware(modelos.GetCompetenciaAll)).Methods("GET")
	router.HandleFunc("/competencia/{id}", ValidateMiddleware(modelos.GetCompetencia)).Methods("GET")
	router.HandleFunc("/competencia/", ValidateMiddleware(modelos.CreateCompetencia)).Methods("POST")
	router.HandleFunc("/competencialist/", ValidateMiddleware(modelos.CreateCompetenciaAll)).Methods("POST")
	router.HandleFunc("/competencia/{id}", ValidateMiddleware(modelos.DeleteCompetencia)).Methods("DELETE")
	router.HandleFunc("/competencia/{id}", ValidateMiddleware(modelos.CreateCompetencia)).Methods("PUT")

	router.HandleFunc("/competenciapatc/", ValidateMiddleware(modelos.GetCompetenciaPaTECAll)).Methods("GET")
	router.HandleFunc("/competenciapaco/", ValidateMiddleware(modelos.GetCompetenciaPaCOMAll)).Methods("GET")
	router.HandleFunc("/competenciapaor/", ValidateMiddleware(modelos.GetCompetenciaPaORGAll)).Methods("GET")

	router.HandleFunc("/competenciapa/", ValidateMiddleware(modelos.GetCompetenciaPaAll)).Methods("GET")
	router.HandleFunc("/competenciapa/{id}", ValidateMiddleware(modelos.GetCompetencia)).Methods("GET")
	router.HandleFunc("/competenciapa/", ValidateMiddleware(modelos.CreateCompetenciaPa)).Methods("POST")
	router.HandleFunc("/competenciapa/{id}", ValidateMiddleware(modelos.DeleteCompetenciaPa)).Methods("DELETE")
	router.HandleFunc("/competenciapa/{id}", ValidateMiddleware(modelos.CreateCompetenciaPa)).Methods("PUT")

	router.HandleFunc("/permissaotipo/", ValidateMiddleware(modelos.GetPermissaoTipoAll)).Methods("GET")
	router.HandleFunc("/permissaotipo/{id}", ValidateMiddleware(modelos.GetPermissaoTipo)).Methods("GET")
	router.HandleFunc("/permissaotipo/", ValidateMiddleware(modelos.CreatePermissaoTipo)).Methods("POST")
	router.HandleFunc("/permissaotipo/{id}", ValidateMiddleware(modelos.DeletePermissaoTipo)).Methods("DELETE")
	router.HandleFunc("/permissaotipo/{id}", ValidateMiddleware(modelos.CreatePermissaoTipo)).Methods("PUT")

	router.HandleFunc("/setor/", ValidateMiddleware(modelos.GetSetorAll)).Methods("GET")
	router.HandleFunc("/setor/{id}", ValidateMiddleware(modelos.GetSetor)).Methods("GET")
	router.HandleFunc("/setor/", ValidateMiddleware(modelos.CreateSetor)).Methods("POST")
	router.HandleFunc("/setor/{id}", ValidateMiddleware(modelos.DeleteSetor)).Methods("DELETE")
	router.HandleFunc("/setor/{id}", ValidateMiddleware(modelos.CreateSetor)).Methods("PUT")

	router.HandleFunc("/sexo/", ValidateMiddleware(modelos.GetSexoAll)).Methods("GET")
	router.HandleFunc("/sexo/{id}", ValidateMiddleware(modelos.GetSexo)).Methods("GET")
	router.HandleFunc("/sexo/", ValidateMiddleware(modelos.CreateSexo)).Methods("POST")
	router.HandleFunc("/sexo/{id}", ValidateMiddleware(modelos.DeleteSexo)).Methods("DELETE")
	router.HandleFunc("/sexo/{id}", ValidateMiddleware(modelos.CreateSexo)).Methods("PUT")

	router.HandleFunc("/termo/", ValidateMiddleware(modelos.GetTermoAll)).Methods("GET")
	router.HandleFunc("/termo/{id}", ValidateMiddleware(modelos.GetTermo)).Methods("GET")
	router.HandleFunc("/termo/", ValidateMiddleware(modelos.CreateTermo)).Methods("POST")
	router.HandleFunc("/termo/{id}", ValidateMiddleware(modelos.DeleteTermo)).Methods("DELETE")
	router.HandleFunc("/termo/{id}", ValidateMiddleware(modelos.CreateTermo)).Methods("PUT")

	router.HandleFunc("/tipocompetencia/", ValidateMiddleware(modelos.GetTipoCompetenciaAll)).Methods("GET")
	router.HandleFunc("/tipocompetencia/{id}", ValidateMiddleware(modelos.GetTipoCompetencia)).Methods("GET")
	router.HandleFunc("/tipocompetencia/", ValidateMiddleware(modelos.CreateTipoCompetencia)).Methods("POST")
	router.HandleFunc("/tipocompetencia/{id}", ValidateMiddleware(modelos.DeleteTipoCompetencia)).Methods("DELETE")
	router.HandleFunc("/tipocompetencia/{id}", ValidateMiddleware(modelos.CreateTipoCompetencia)).Methods("PUT")

	router.HandleFunc("/usuario/", ValidateMiddleware(modelos.CreateUsuario)).Methods("POST")
	router.HandleFunc("/usuario/", ValidateMiddleware(modelos.GetUsuarioAll)).Methods("GET")
	router.HandleFunc("/usuario/{id}", ValidateMiddleware(modelos.CreateUsuario)).Methods("PUT")
	router.HandleFunc("/usuario/{id}", ValidateMiddleware(modelos.DeleteUsuario)).Methods("DELETE")

	router.HandleFunc("/funcionario/", ValidateMiddleware(modelos.CreateFuncionario)).Methods("POST")
	router.HandleFunc("/funcionario/", ValidateMiddleware(modelos.GetFuncionarioAll)).Methods("GET")
	router.HandleFunc("/funcionariopg/{page}", ValidateMiddleware(modelos.GetFuncionarioPGAll)).Methods("GET")
	router.HandleFunc("/funcionario/{id}", ValidateMiddleware(modelos.CreateFuncionario)).Methods("PUT")
	router.HandleFunc("/funcionario/{id}", ValidateMiddleware(modelos.DeleteFuncionario)).Methods("DELETE")

	router.HandleFunc("/gestor/", ValidateMiddleware(modelos.CreateGestor)).Methods("POST")
	router.HandleFunc("/gestor/", ValidateMiddleware(modelos.GetGestorAll)).Methods("GET")
	router.HandleFunc("/gestor/{id}", ValidateMiddleware(modelos.CreateGestor)).Methods("PUT")
	router.HandleFunc("/gestor/{id}", ValidateMiddleware(modelos.DeleteGestor)).Methods("DELETE")

	//router.HandleFunc("/sdc/{obra}/{datainicio}/{datafim}", ValidateMiddleware(modelos.GetSDCConsulta)).Methods("GET")

	router.HandleFunc("/tarefas/", ValidateMiddleware(modelos.GetTarefasAll)).Methods("GET")
	router.HandleFunc("/tarefas/", ValidateMiddleware(modelos.CreateTarefas)).Methods("POST")
	router.HandleFunc("/tarefas/{id}", ValidateMiddleware(modelos.CreateTarefas)).Methods("PUT")
	router.HandleFunc("/tarefas/{id}", ValidateMiddleware(modelos.DeleteTarefas)).Methods("DELETE")

	router.HandleFunc("/validateToken", ValidateMiddleware(modelos.ValidateToken)).Methods("POST")

	handler := c.Handler(router)
	if os.Getenv("PORT") == "" {

		http.ListenAndServe(":8080", handler)
	} else {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Println("Port was defined but could not be parsed.")
			os.Exit(1)
		}
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
	}
}

func exec(db *sql.DB, sql string) sql.Result {
	// abri uma conexão com o banco
	result, err := db.Exec(sql)
	if err != nil {
		fmt.Println("Teste erro", err)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Erro na conexao com o Banco de Dados")
		panic(err)

	}
	return result
}

func criarBanco() {

	db, err := sql.Open("postgres", "postgres://postgres:andre110407@localhost/cerh?sslmode=disable")
	//db, err := sql.Open("postgres", "postgres://postgres:andre110407@localhost/sisalicerce?sslmode=disable")
	//DB, err = sql.Open("postgres", "postgres://postgres:Andre110407@localhost/lordchicken?sslmode=disable")
	if err != nil {
		panic(err)
	}
	// fecha antes do final de main
	defer db.Close()

	exec(db, `create table CompetenciaPa (
		id serial PRIMARY KEY NOT NULL,
		datacadastro timestamp without time zone,
		titulo varchar(80),
		descricao varchar(80),
		ativo BOOLEAN NOT NULL DEFAULT FALSE,
		id_tipocompetencia integer,
		FOREIGN KEY (id_tipocompetencia) REFERENCES tipocompetencia (id),

	 	id_gestor integer,
		FOREIGN KEY (id_gestor) REFERENCES gestor (id),
		id_empresa integer,
		FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table questionarioavaliacao (
	// 	id serial PRIMARY KEY NOT NULL,
	// 	nome varchar(80),
	// 	datacadastro timestamp without time zone,
	// 	dataquestionario timestamp without time zone,
	// 	dataresposta timestamp without time zone,
	// 	auditoria varchar(80),
	// 	respondida BOOLEAN NOT NULL DEFAULT FALSE,
	// 	pontuacaofinal integer,
	// 	peso integer,
	// 	valorpesoavaliacao integer,

	// 	id_avaliacao integer,
	// 	FOREIGN KEY (id_avaliacao) REFERENCES avaliacao (id),

	//  	id_funcionarioavaliado integer,
	// 	FOREIGN KEY (id_funcionario) REFERENCES funcionario (id),

	// 	id_funcionarioavaliador integer,
	//  	FOREIGN KEY (id_funcionario) REFERENCES funcionario (id),

	// 	id_empresa integer,
	// 	FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table avaliacao (
	// 	id serial PRIMARY KEY NOT NULL,
	// 	datacadastro timestamp without time zone,
	// 	login varchar(80),

	// 	valortotalavaliacoes real,

	// 	id_funcionarioavaliado integer,
	// 	FOREIGN KEY (id_funcionario) REFERENCES funcionario (id),

	// 	id_funcionarioavaliador1 integer,
	// 	FOREIGN KEY (id_funcionario) REFERENCES funcionario (id),
	// 	avaliacaoefetuada1 BOOLEAN NOT NULL DEFAULT FALSE,
	// 	valortotalavaliacao1 real,
	// 	valorpesoavaliacao1 real,
	// 	dataavaliacaoefetuada1 timestamp without time zone,

	// 	id_funcionarioavaliador2 integer,
	// 	FOREIGN KEY (id_funcionario) REFERENCES funcionario (id),
	// 	avaliacaoefetuada2 BOOLEAN NOT NULL DEFAULT FALSE,
	// 	valortotalavaliacao2 real,
	// 	valorpesoavaliacao2 real,
	// 	dataavaliacaoefetuada2 timestamp without time zone,

	// 	id_funcionarioavaliador3 integer,
	// 	FOREIGN KEY (id_funcionario) REFERENCES funcionario (id),
	// 	avaliacaoefetuada3 BOOLEAN NOT NULL DEFAULT FALSE,
	// 	valortotalavaliacao3 real,
	// 	valorpesoavaliacao3 real,
	// 	dataavaliacaoefetuada3 timestamp without time zone,

	// 	id_funcionarioavaliador44 integer,
	// 	FOREIGN KEY (id_funcionario) REFERENCES funcionario (id),
	// 	avaliacaoefetuada4 BOOLEAN NOT NULL DEFAULT FALSE,
	// 	valortotalavaliacao4 real,
	// 	valorpesoavaliacao4 real,
	// 	dataavaliacaoefetuada4 timestamp without time zone,

	// 	id_gestor integer,
	// 	FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// 	id_empresa integer,
	// 	FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table SEXO (
	// 	id serial PRIMARY KEY NOT NULL,
	// 	nome varchar(80),
	// 	datacadastro timestamp without time zone,
	// 	id_gestor integer,
	// 	FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// 	id_empresa integer,
	// 	FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table Cargo (
	// 	id serial PRIMARY KEY NOT NULL,
	// 	nome varchar(80),
	// 	datacadastro timestamp without time zone,
	// 	id_gestor integer,
	// 	FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// 	id_empresa integer,
	// 	FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table departamento (
	// 	id serial PRIMARY KEY NOT NULL,
	// 	nome varchar(80),
	// 	datacadastro timestamp without time zone,
	// 	id_gestor integer,
	// 	FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// 	id_empresa integer,
	// 	FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table funcao (
	// 	id serial PRIMARY KEY NOT NULL,
	// 	nome varchar(80),
	// 	datacadastro timestamp without time zone,
	// 	id_gestor integer,
	// 	FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// 	id_empresa integer,
	// 	FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table grauinstrucao (
	// 	id serial PRIMARY KEY NOT NULL,
	// 	nome varchar(80),
	// 	datacadastro timestamp without time zone,
	// 	id_gestor integer,
	// 	FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// 	id_empresa integer,
	// 	FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table observacaodois (
	// id serial PRIMARY KEY NOT NULL,
	// datacadastro timestamp without time zone,
	// obs varchar(80),
	// melhorias varchar(80),
	// id_avaliacaodois integer,
	// FOREIGN KEY (id_avaliacaodois) REFERENCES avaliacaodois (id),
	// id_gestor integer,
	// FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// id_empresa integer,
	// FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table observacao (
	// id serial PRIMARY KEY NOT NULL,
	// datacadastro timestamp without time zone,
	// obs varchar(80),
	// melhorias varchar(80),
	// id_avaliacao integer,
	// FOREIGN KEY (id_avaliacao) REFERENCES avaliacao (id),
	// id_gestor integer,
	// FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// id_empresa integer,
	// FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	//  exec(db, `create table peso (
	// 	id serial PRIMARY KEY NOT NULL,

	//  	datacadastro timestamp without time zone,
	// 	 pesosuperior integer,
	// 	 pesoparceiro integer,
	// 	 pesosubordinado integer,
	// 	 pesoautoavaliacao integer,
	// 	id_gestor integer,
	//     FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	//  	id_empresa integer,
	//     FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table setor (
	// 	id serial PRIMARY KEY NOT NULL,
	//    nome varchar(80),
	// 	datacadastro timestamp without time zone,
	//    id_gestor integer,
	//    FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// 	id_empresa integer,
	//    FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table Sexo (
	//  	id serial PRIMARY KEY NOT NULL,
	//     nome varchar(80),
	//  	datacadastro timestamp without time zone,
	// 	id_gestor integer,
	//     FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	//  	id_empresa integer,
	//     FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table Tarefas (
	// 	id serial PRIMARY KEY NOT NULL,
	//     descricao varchar(80),
	// 	datacadastro timestamp without time zone,
	// 	concluido BOOLEAN NOT NULL DEFAULT FALSE,
	//     id_gestor integer,
	//     FOREIGN KEY (id_gestor) REFERENCES gestao (id),
	// 	id_empresa integer,
	//     FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table Termo (
	// 	id serial PRIMARY KEY NOT NULL,
	//     nome varchar(80),
	//  datacadastro timestamp without time zone,
	//   abreviacao  varchar(80),
	//   id_gestao integer,
	//   FOREIGN KEY (id_gestao) REFERENCES gestao (id),
	//   id_avaliacao integer,
	//   FOREIGN KEY (id_avaliacao) REFERENCES avaliacao (id),
	//   id_empresa integer,
	//   FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table tipocompetencia (
	//   			id serial PRIMARY KEY NOT NULL,
	//    			nome varchar(80),
	//    			datacadastro timestamp without time zone,
	// 			abreviacao  varchar(80),
	// 			id_gestor integer,
	//     		FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	//     		id_empresa integer,
	// 		    FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, `create table setor (
	//   			id serial PRIMARY KEY NOT NULL,
	//   			nome varchar(80),
	//   			datacadastro timestamp without time zone,
	// 			id_gestor integer,
	// 			FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// 			id_empresa integer,
	// 			   FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, "drop table if exists gestor")
	// exec(db, `create table gestor (
	//   			id serial PRIMARY KEY NOT NULL,
	// 			nome varchar(80),
	// 			nomegestor varchar(80),
	//   			telefone varchar(80),
	//    			login varchar(80),
	//    			senha bytea,
	//  			datacadastro timestamp without time zone,
	// 			auditoria varchar(80),
	// 			primeiroacesso BOOLEAN NOT NULL DEFAULT FALSE,
	// 			ativo BOOLEAN NOT NULL DEFAULT FALSE,
	// 			dataativo  timestamp without time zone,
	//   			validade BOOLEAN NOT NULL DEFAULT FALSE,
	//  			datavalidade  timestamp without time zone,
	//  			id_empresa integer,
	//        		FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, "drop table if exists competencia")
	// exec(db, `create table competencia (
	// 		id serial PRIMARY KEY NOT NULL,
	//		nome varchar(80),
	//		titulo varchar(80),
	//		datacadastro timestamp without time zone,
	//		peso integer,
	//		ativo BOOLEAN NOT NULL DEFAULT FALSE,
	//		abreviacao varchar(80),
	//      descricao varchar(80),
	//  	id_tipocompetencia integer,
	//  	FOREIGN KEY (id_tipocompetencia) REFERENCES tipocompetencia (id),
	//  	id_cargo integer,
	//      FOREIGN KEY (id_cargo) REFERENCES cargo (id),
	// 		id_gestor integer,
	//      FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	//  	id_empresa integer,
	//      FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// exec(db, "drop table if exists usuario")
	// exec(db, `create table usuario (
	// 		id serial PRIMARY KEY NOT NULL,
	// 		nome varchar(80),
	// 			telefone varchar(80),
	//    			login varchar(80),
	//    			senha bytea,
	//  			datacadastro timestamp without time zone,
	//  			dataadmissao timestamp without time zone,
	//  			datanascimento timestamp without time zone,
	//   			auditoria varchar(80),
	//   			primeiroacesso BOOLEAN NOT NULL DEFAULT FALSE,
	//    			ativo BOOLEAN NOT NULL DEFAULT FALSE,
	//    			dataativo  timestamp without time zone,
	//   			validade BOOLEAN NOT NULL DEFAULT FALSE,
	//  			datavalidade  timestamp without time zone,
	//  			id_cargo integer,
	//  			FOREIGN KEY (id_cargo) REFERENCES cargo (id),
	//  			id_setor integer,
	//        		FOREIGN KEY (id_setor) REFERENCES setor (id),
	// 			id_gestor integer,
	//        		FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	//  			id_empresa integer,
	//        		FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// // 			fmt.Println("2 tabelmas   ok")
	// // 			exec(db, "drop table if exists tarefas")
	// // 			exec(db, `create table tarefas (
	// // 	id serial PRIMARY KEY NOT NULL,
	// // 	descricao varchar(200),
	// // 	datacadastro timestamp without time zone,
	// // 	concluido BOOLEAN NOT NULL DEFAULT FALSE,
	// // 	id_usuario integer,
	// // 	FOREIGN KEY (id_usuario) REFERENCES usuario (id),
	// // 	id_empresa integer,
	// // 	FOREIGN KEY (id_empresa) REFERENCES empresa (id))

	// // fmt.Println("criou 1 Empresa ok")
	// // exec(db, "drop table if exists empresa")
	// // exec(db, `create table empresa (
	// 			id serial PRIMARY KEY NOT NULL,
	// 			razaosocial varchar(80),
	// 			dataCadastro  varchar(80),
	// 			email varchar(80),
	// 			emailsecu varchar(80),
	// 			telefone varchar(100),
	// 			telefoneSecu varchar(100),
	// 			cpf varchar(100),
	// 			cnpj varchar(100),
	// 			pessoaFisica BOOLEAN NOT NULL DEFAULT TRUE,
	// 			acesso BOOLEAN NOT NULL DEFAULT TRUE,
	// 			obsPessoa varchar(200))`)

	// fmt.Println("2 enderecoempresa   ok")
	// exec(db, "drop table if exists enderecoempresa")
	// exec(db, `create table enderecoempresa (
	// 	   		id serial PRIMARY KEY NOT NULL,
	// 	   		Cep varchar(80),
	// 			Endereco varchar(80),
	// 			Bairro varchar(80),
	// 			Cidade varchar(80),
	// 			Complemento varchar(80),
	// 			Numero integer,
	// 			Uf varchar(80),
	// 			Ddd varchar(80),
	// 			Unidade varchar(80),
	// 			Ibge varchar(80),
	// 			id_empresa integer,
	// 			FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// fmt.Println("4, permissaotipo ok ")
	// exec(db, "drop table if exists permissaotipo")
	// exec(db, `create table permissaotipo (
	// 		    id serial PRIMARY KEY NOT NULL,
	// 			nome varchar(80),
	// 			descricao varchar(80),
	// 			dataCadastro varchar(80),
	// 			id_usuario integer,
	// 			FOREIGN KEY (id_usuario) REFERENCES usuario (id),
	// 			id_empresa integer,
	// 			FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// fmt.Println("5, permissao ok ")
	// exec(db, "drop table if exists permissao")
	// exec(db, `create table permissao (
	// 			id serial PRIMARY KEY NOT NULL,
	// 			nome varchar(80),
	// 			descricao varchar(80),
	// 			dataCadastro varchar(80),
	// 			id_permissaotipo integer,
	// 			FOREIGN KEY (id_permissaotipo) REFERENCES permissaotipo (id),
	// 			id_usuario integer,
	// 			FOREIGN KEY (id_usuario) REFERENCES usuario (id),
	// 			id_empresa integer,
	// 			FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

	// fmt.Println("5, funcionario ok ")
	// exec(db, "drop table if exists funcionario")
	// exec(db, `create table funcionario (
	//  		id serial PRIMARY KEY NOT NULL,
	// 			nome varchar(80),
	// 			telefone varchar(80),
	// 			login varchar(80),
	// 			senha bytea,
	// 			datacadastro timestamp without time zone,
	// 			dataadmissao timestamp without time zone,
	// 			datanascimento timestamp without time zone,
	// 			auditoria varchar(80),
	// 			codigo varchar(20),
	// 			primeiroacesso BOOLEAN NOT NULL DEFAULT FALSE,
	// 			ativo BOOLEAN NOT NULL DEFAULT FALSE,
	// 			dataativo timestamp without time zone,
	// 			validado BOOLEAN NOT NULL DEFAULT FALSE,
	// 			datavalidado timestamp without time zone,
	// 			 id_gestor integer,
	//  			FOREIGN KEY (id_gestor) REFERENCES gestor (id),
	// 			id_sexo integer,
	//  			FOREIGN KEY (id_sexo) REFERENCES sexo (id),
	//  			id_cargo integer,
	//  			FOREIGN KEY (id_cargo) REFERENCES cargo (id),
	//  			id_setor integer,
	//  			FOREIGN KEY (id_setor) REFERENCES setor (id),
	//  			id_empresa integer,
	// 			FOREIGN KEY (id_empresa) REFERENCES empresa (id))`)

}

func sendEmail() {
	//Criamos um slice do tipo string do tamanho máximo de 1 para receber nosso e-mail destinatário.
	recipients := make([]string, 1)
	recipients[0] = "andreggp@gmail.com"

	err := smtp.SendMail(
		/* endereço do servidor de SMTP */ "smtp.gmail.com:25",
		/* mecanismo de autenticação*/ smtp.PlainAuth("", "andreggp@gmail.com", "andre110407", "smtp.gmail.com"),
		/* e-mail de origem */ "andreggp@gmail.com",
		/*Mensagem no RFC 822-style*/ recipients,
		/* Corpo da mensagem */ []byte("Subject:Olá!\n\n Olá Fulano. Tudo de bom com Go!"))
	if err != nil {
		log.Fatal(err)
	}
}

type smtpServer struct {
	host string
	port string
}

// serverName URI to smtp server
func (s *smtpServer) serverName() string {
	return s.host + ":" + s.port
}

// Address URI to smtp server.
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}
func sendEmailTeste() {
	// Sender data.
	from := "andreggp@gmail.com"
	password := "andre110407"
	// Receiver email address.
	to := []string{
		"andreggp@gmail.com",
		"andreggp@gmail.com",
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	message := []byte("This is a really unimaginative message, I know.")
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}

//ValidateMiddleware validade resuisições com token
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authorizationHeader := req.Header.Get("authorization")

		// authorizationHeader1 := req.Header.Get("headers")
		// fmt.Println("Hearder", req.Header.Get("headers"))
		// fmt.Println("TESTE", authorizationHeader1)

		// decoder := json.NewDecoder(req.Body)

		// payload := make(map[string]interface{})
		// err := decoder.Decode(&payload)
		// fmt.Println(err)
		// var token string
		// var Bearer string

		// fmt.Println("payload ", payload)

		// for i, v := range payload {
		// 	fmt.Println(i, "=", v)

		// 	switch i {
		// 	case "Bearer":
		// 		token, _ = v.(string)
		// 		Bearer = i
		// 		continue
		// 	}
		// }

		// fmt.Println("Token ", token)
		// fmt.Println("Bearer ", Bearer)

		// // /var jwtToken JwtToken
		// // erro := json.NewDecoder(req.Body).Decode(&jwtToken)

		// // fmt.Println("jwtToken: ", jwtToken)
		// // fmt.Println("erro: ", erro)
		// // fmt.Println("Entrou aqui jwtToken : ", jwtToken)
		// // fmt.Println("Entrou aqui err      : ", err)

		// //authorizationHeader := req.Header.Get("authorization")

		// authorizationHeader := Bearer + " " + token

		// fmt.Println("authorizationHeader :", authorizationHeader)

		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")

			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Token invalido")
					}
					return controler.JwtSecretKey, nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(Exception{Message: error.Error()})
					return
				}
				if token.Valid {

					//return res.status(200).send({ valid: !err })

					//pega os dos empresa e usuario dentro do token
					empresa := fmt.Sprintf("%v", token.Claims.(jwt.MapClaims)["empresa"])
					gestor := fmt.Sprintf("%v", token.Claims.(jwt.MapClaims)["gestor"])

					//seta as informações de empresa e usuario dentro da requisição
					req.Header.Set("empresa", empresa)
					req.Header.Set("gestor", gestor)
					//req.Header.Set({"Valid": "true"})

					//json.NewEncoder(w).Encode()

					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(Exception{Message: "Token invalido"})
				}
			}
		} else {
			fmt.Println("Erro ao extrair jwt")
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}
	})
}

//JwtExtract extract
func JwtExtract(r *http.Request) (map[string]interface{}, error) {
	tokenString := ExtractBearerToken(r)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return controler.JwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}
	return claims, nil
}

//ExtractBearerToken token
func ExtractBearerToken(r *http.Request) string {
	headerAuthorization := r.Header.Get("Authorization")
	bearerToken := strings.Split(headerAuthorization, " ")
	return html.EscapeString(bearerToken[1])
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	teste := "1 André e Lindo"
	p1Json, _ := json.Marshal(teste)
	w.Write(p1Json)
}

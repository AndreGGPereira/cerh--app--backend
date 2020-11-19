package persistencia

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//DB teste
var DB *sql.DB

func init() {
	fmt.Println("Tentou conexão")

	var err error
	//DB, err = sql.Open("postgres", "postgres://tlaqkgxcydmewb:e8c9fec844a38bafe2baeefad88523f5e4bd39f0253b97e452f514acb76be592@ec2-174-129-242-183.compute-1.amazonaws.com:5432/d6clrav5662k3i")

	//heroku atual
	//DB, err = sql.Open("postgres", "postgres://mgulakptkfujrh:0b2e4e2d38ed347e722d162ae00782b69c4181ae831bf3516746d40f4481b973@ec2-174-129-227-80.compute-1.amazonaws.com:5432/d6ehr2l533fta7")
	//DB, err = sql.Open("postgres", "postgres://postgres:Andre110407@localhost/lordchicken?sslmode=disable")
	//local
	//SisAlicerce postgres://tgzsrbubcehoem:f97506a59da9f77b8dd6097fda5dfa2331635fb9b0f1bcba1cdab82796c53d49@ec2-3-91-112-166.compute-1.amazonaws.com:5432/dcorodteg3avdo

	//Local DB, err = sql.Open("postgres", "postgres://postgres:andre110407@localhost/sisalicerce?sslmode=disable")

	//DB, err = sql.Open("postgres", "postgres://tgzsrbubcehoem:f97506a59da9f77b8dd6097fda5dfa2331635fb9b0f1bcba1cdab82796c53d49@ec2-3-91-112-166.compute-1.amazonaws.com:5432/dcorodteg3avdo")
	DB, err = sql.Open("postgres", "postgres://postgres:Andre110407@localhost/cerh?sslmode=disable")
	fmt.Println("Testanto a conexao banco de dados - persistencia/conexaoDB")

	if err != nil {
		fmt.Println("Erro ao acessar o banco")
		fmt.Println(err)
		//panic(err)
	}
	if err = DB.Ping(); err != nil {
		fmt.Println("Erro na conexao com o Banco de Dados")
		fmt.Println(err)

	}
	fmt.Println("Conexão Estabelecida")
}

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

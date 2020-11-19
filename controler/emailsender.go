package controler

import (
	"fmt"
	"log"
	"net/smtp"
)

//Send email
func Send(body string) {
	from := "andreggp@gmail.com"
	pass := "andre110407"
	to := "andreggp@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}

//EnviarEmail teste
func EnviarEmail(email, token string) {

	//Criamos um slice do tipo string do tamanho máximo de 1 para receber nosso e-mail destinatário.
	msg := " token : " + token

	recipients := make([]string, 1)
	recipients[0] = "andreggp@gmail.com"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>    <a href="http:localhost:3000/` + token + `>Acesso aqui para validar seu email</a>
		</h1>
	</body>
	</html>
`)

	fmt.Println(msg)
	fmt.Println(str)
	err := smtp.SendMail(
		/* endereço do servidor de SMTP */ "smtp.gmail.com:25",
		/* mecanismo de autenticação*/ smtp.PlainAuth("", "andreggp@gmail.com", "andre110407", "smtp.gmail.com"),
		/* e-mail de origem */ "andreggp@gmail.com",
		/*Mensagem no RFC 822-style*/ recipients,
		/* Corpo da mensagem */ []byte(str))
	if err != nil {
		log.Fatal(err)
	}

}

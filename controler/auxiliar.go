package controler

import (
	"math/rand"
	"net/http"
)

//PegarIDEmpresaIDUsuario fornece os id dos itens oriundos da requisição
func PegarIDEmpresaIDUsuario(req *http.Request) (idempresa int, idusuario int) {

	idempresa = StringForInt(req.Header.Get("empresa"))
	idusuario = StringForInt(req.Header.Get("usuario"))

	return idempresa, idusuario
}

//PegarIDEmpresaIDFuncionario dados esto
func PegarIDEmpresaIDFuncionario(req *http.Request) (idempresa int, idusuario int) {

	idempresa = StringForInt(req.Header.Get("empresa"))
	idusuario = StringForInt(req.Header.Get("funcionario"))

	return idempresa, idusuario
}

//PegarIDEmpresaIDGestor fornece os id dos itens oriundos da requisição
func PegarIDEmpresaIDGestor(req *http.Request) (idempresa int, idgestor int) {

	idempresa = StringForInt(req.Header.Get("empresa"))
	idgestor = StringForInt(req.Header.Get("gestor"))
	return idempresa, idgestor
}

var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// generates a random string of fixed size
func srand(size int) string {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

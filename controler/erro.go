package controler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Error struct
type Error struct {
	Err        json.RawMessage `json:"error"`
	StatusCode int             `json:"statuscode"`
}

//New novo erro
func New(err string) (apierr Error) {
	ret := struct {
		Error  string `json:"error"`
		Status string `json:"status"`
	}{
		Error:  err,
		Status: "erro",
	}
	b, e := json.MarshalIndent(ret, "", "\t")
	if e != nil {
		log.Fatal(e)
	}
	apierr.Err = b
	apierr.StatusCode = http.StatusInternalServerError
	return
}

//NewRawJSON formatação
func NewRawJSON(rawJSON json.RawMessage, statusCode int) (apierr Error) {
	apierr.Err = rawJSON
	apierr.StatusCode = statusCode
	return
}

//Error retorna o erro
func (e Error) Error() string {
	return string(e.Err)
}

//Writer  erro
func (e Error) Write(w http.ResponseWriter) {
	log.Panic(e.Error())
	//log.Error(e.Error())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)
	fmt.Fprintln(w, string(e.Err))
}

//ErrorReport erro reportado
func ErrorReport(w http.ResponseWriter, err error, statusCode int) {
	if e, ok := err.(Error); ok {
		e.StatusCode = statusCode
		e.Write(w)
		return
	}
	e := New(err.Error())
	e.StatusCode = statusCode
	e.Write(w)
}

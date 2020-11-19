package controler

import "encoding/json"

//TokenErro dados do token
type TokenErro struct {
	Err        json.RawMessage `json:"error"`
	StatusCode int             `json:"statuscode"`
}

//JwtToken struct
type JwtToken struct {
	Token  string `json:"token"`
	Status int    `json:"status"`
	Login  string `json:"login"`
	Nome   string `json:"nome"`
}

//JwtTokenGestor struct
type JwtTokenGestor struct {
	Token   string `json:"token"`
	Status  int    `json:"status"`
	Gestor  string `json:"gestor"`
	Usuario string `json:"usuario"`
	Nome    string `json:"nome"`
}

//JwtTokenValid token
type JwtTokenValid struct {
	Valid bool `json:"valid"`
}

package controler

//Message envia mensagens para json
type Message struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

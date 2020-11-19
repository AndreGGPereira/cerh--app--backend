package modelos

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

//BrCep struct para Cep
type BrCep struct {
	Cep         string `json:"cep,omitempty"`
	Endereco    string `json:"endereco,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Cidade      string `json:"cidade,omitempty"`
	Uf          string `json:"uf,omitempty"`
	Latitude    string `json:"latitude,omitempty"`
	Longitude   string `json:"longitude,omitempty"`
	Ddd         string `json:"ddd,omitempty"`
	Unidade     string `json:"unidade,omitempty"`
	Ibge        string `json:"ibge,omitempty"`
}

//BuscaCep consulta de CEP
func BuscaCep(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// decoder := json.NewDecoder(req.Body)

	// var obj BrCep
	// // err := decoder.Decode(&obj)
	// // if err != nil {
	// // 	json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
	// // 	fmt.Println("Erro", err)
	// // }

	// payload := make(map[string]interface{})
	// err := decoder.Decode(&payload)
	// fmt.Println(err)

	// fmt.Println("payload", payload)

	// for i, v := range payload {
	// 	fmt.Println(i, "=", v)
	// }

	// fmt.Println("CEP", obj)
	// fmt.Println("decoder", decoder)
	//cep := "01306020"

	var obj BrCep
	var err error
	vars := mux.Vars(req)

	fmt.Println("Dados do cel ", vars)
	//obj.Cep, err = strconv.Atoi(vars["id"])
	obj.Cep, _ = vars["cep"]

	if err != nil {
		json.NewEncoder(w).Encode("Erro ao inserir o identificador")
	}

	cepSeguro := url.QueryEscape(obj.Cep)

	url := fmt.Sprintf("https://brcep-dlfeappmhe.now.sh/%s/json", cepSeguro)

	req1, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	resp, err := client.Do(req1)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()
	var resultado BrCep

	if err := json.NewDecoder(resp.Body).Decode(&resultado); err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(&resultado)
}

//GetCep consulta de CEP
func GetCep(cep string) BrCep {

	fmt.Println("Entrou aqui cep")

	var obj BrCep
	obj.Cep = cep
	fmt.Println("CEP", obj.Cep)

	//cep := "01306020"
	cepSeguro := url.QueryEscape(obj.Cep)

	url := fmt.Sprintf("https://brcep-dlfeappmhe.now.sh/%s/json", cepSeguro)

	req1, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	resp, err := client.Do(req1)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	defer resp.Body.Close()
	var resultado BrCep

	if err := json.NewDecoder(resp.Body).Decode(&resultado); err != nil {
		log.Println(err)
	}

	return resultado
}

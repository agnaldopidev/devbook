package resposta

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErroAPI struct {
	Erro string `json:"erro"`
}

// Retorna resposta em json
func JSON(w http.ResponseWriter, statusCod int, dados interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCod)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Tratar resposta com code 400 ou mais
func TratarStatusCod(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}

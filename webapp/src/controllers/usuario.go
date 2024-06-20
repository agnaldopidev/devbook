package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/resposta"
)

func Cadastrar(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Println(r.FormValue("nome"))

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		resposta.JSON(w, http.StatusBadRequest, resposta.ErroAPI{Erro: erro.Error()})
		return
	}

	fmt.Println("usuario", bytes.NewBuffer(usuario))
	url := fmt.Sprintf("%s/login", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))

	if erro != nil {
		resposta.JSON(w, http.StatusInternalServerError, resposta.ErroAPI{Erro: erro.Error()})
	}

	if response.StatusCode >= 400 {
		resposta.TratarStatusCod(w, response)
		return
	}

	defer response.Body.Close()
	resposta.JSON(w, response.StatusCode, nil)
}

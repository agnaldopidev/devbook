package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookie"
	"webapp/src/modelo"
	"webapp/src/resposta"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		resposta.JSON(w, http.StatusBadRequest, resposta.ErroAPI{Erro: erro.Error()})
		return
	}
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

	if response.StatusCode >= 400 {
		resposta.TratarStatusCod(w, response)
		return
	}

	var autenticao modelo.Autenticacao

	if erro := json.NewDecoder(response.Body).Decode(&autenticao); erro != nil {
		resposta.JSON(w, http.StatusUnprocessableEntity, resposta.ErroAPI{Erro: erro.Error()})
		return
	}

	if erro = cookie.SalvarNoCookie(w, autenticao); erro != nil {
		resposta.JSON(w, http.StatusUnprocessableEntity, resposta.ErroAPI{Erro: erro.Error()})
		return
	}
	resposta.JSON(w, response.StatusCode, nil)
}

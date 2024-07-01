package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicao"
	"webapp/src/resposta"
)

func Publicar(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		resposta.JSON(w, http.StatusBadRequest, resposta.ErroAPI{Erro: erro.Error()})
		return
	}

	fmt.Println("\n\t\tpublicacao\n", bytes.NewBuffer(publicacao))
	url := fmt.Sprintf("%s/publicacao", config.APIURL)
	//response, erro := http.Post(url, "application/json", bytes.NewBuffer(publicacao))
	response, erro := requisicao.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
	if erro != nil {
		resposta.JSON(w, http.StatusInternalServerError, resposta.ErroAPI{Erro: erro.Error()})
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		resposta.TratarStatusCod(w, response)
		return
	}

	defer response.Body.Close()
	resposta.JSON(w, response.StatusCode, nil)
}

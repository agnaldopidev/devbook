package controllers

import (
	//"encoding/json"
	//"fmt"
	"net/http"
	//"webapp/src/config"
	//"webapp/src/modelo"
	//"webapp/src/requisicao"
	//"webapp/src/resposta"
	"webapp/src/utils"
)

// endpoid login
func CarregaPagLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// endpoid login
func CarregaPagCadastro(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "cadastro.html", nil)
}

// endpoid login
func CarregaPagHome(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}
/*
// Carrega pagina home
func CarregaPagHome(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacao", config.APIURL)
	response, erro := requisicao.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		resposta.JSON(w, http.StatusInternalServerError, erro)
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		resposta.TratarStatusCod(w, response)
		return
	}
	var publicacao []modelo.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacao); erro != nil {
		resposta.JSON(w, http.StatusUnprocessableEntity, resposta.ErroAPI{Erro: erro.Error()})
	}

	utils.ExecuteTemplate(w, "home.html", struct{ Publicacaoes []modelo.Publicacao }{
		Publicacaoes: publicacao,
	})

}
	*/
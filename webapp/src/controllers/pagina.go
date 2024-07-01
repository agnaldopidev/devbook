package controllers

import (
	//"encoding/json"
	//"fmt"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//"webapp/src/config"
	//"webapp/src/modelo"
	//"webapp/src/requisicao"
	//"webapp/src/resposta"

	"webapp/src/config"
	"webapp/src/cookie"
	"webapp/src/modelo"
	"webapp/src/requisicao"
	"webapp/src/resposta"
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
// func CarregaPagHome(w http.ResponseWriter, r *http.Request) {
// 	utils.ExecuteTemplate(w, "home.html", nil)
// }

// Carrega pagina home
func CarregaPagHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n")
	//fmt.Println("Carrega 01")
	url := fmt.Sprintf("%s/publicacao", config.APIURL)
	fmt.Println("Carrega 01", url)
	response, erro := requisicao.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		resposta.JSON(w, http.StatusInternalServerError, erro)
		return
	}
	fmt.Println("Carrega 02", response.StatusCode)
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		resposta.TratarStatusCod(w, response)
		return
	}
	fmt.Println("Carrega 03")
	var publicacacoes []modelo.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacacoes); erro != nil {
		resposta.JSON(w, http.StatusUnprocessableEntity, resposta.ErroAPI{Erro: erro.Error()})
	}

	cookies, _ := cookie.LerValoresDoCookie(r)
	usuarioID, _ := strconv.ParseUint(cookies["id"], 10, 64)
	fmt.Println(publicacacoes)
	fmt.Println(usuarioID)

	utils.ExecuteTemplate(w, "home.html", struct {
		Publicacacoes []modelo.Publicacao
		UsuarioID     uint64
	}{
		Publicacacoes: publicacacoes,
		UsuarioID:     usuarioID,
	})
}

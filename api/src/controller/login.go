package controller

import (
	"api/src/autenticacao"
	"api/src/modelo"
	"api/src/resposta"
	"api/src/seguranca"
	"api/src/service"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user modelo.User
	if erro = json.Unmarshal(corpoRequest, &user); erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userDb, erro := service.FindUserByEmail(user.Email)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	//fmt.Println("\nuser", userDb.Email, userDb.Senha)
	if userDb.Email == "" {
		resposta.Erro(w, http.StatusBadRequest, errors.New("Usuário nao existe"))
		return
	}

	//fmt.Println("user db", userDb.Email)
	if erro = seguranca.Compare(userDb.Senha, user.Senha); erro != nil {
		fmt.Println("usuario nao autorizado senha db", userDb.Senha)
		fmt.Println("usuario nao autorizado senha user", user.Senha)
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(userDb)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	id := strconv.FormatUint(userDb.ID, 10)

	auth := modelo.Autenticacao{ID: id, Token: token}

	resposta.JSON(w, http.StatusOK, auth)
}

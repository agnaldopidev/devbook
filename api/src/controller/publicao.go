package controller

import (
	"api/src/autenticacao"
	"api/src/modelo"
	"api/src/resposta"
	"api/src/service"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Salvar publicacao
func SalvarPublicacao(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	var publicacao modelo.Publicacao
	if erro = json.Unmarshal(corpoRequest, &publicacao); erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
	}

	publicacao, erro = service.SalvarPublicacao(publicacao)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	resposta.JSON(w, http.StatusOK, publicacao)
}

// buscar publicacao
func BuscarPucacao(w http.ResponseWriter, r *http.Request) {
	usuarioLogadoId, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	lista, erro := service.BuscarPucacao(usuarioLogadoId)

	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
	}
	if lista != nil {
		resposta.JSON(w, http.StatusOK, lista)
	}

	resposta.JSON(w, http.StatusBadRequest, nil)
}

// buscar publicacao por id
func BuscarPublicacaoPorId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}

	lista, erro := service.BuscarPublicacaoPorId(userId)

	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
	}
	if lista != nil {
		resposta.JSON(w, http.StatusOK, lista)
	}

	resposta.JSON(w, http.StatusBadRequest, nil)

}

// atualiza publicacao
func UpdatePublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}

	var publicacao modelo.Publicacao
	if erro = json.Unmarshal(corpoRequest, &publicacao); erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
	}

	if erro = service.Permissoes(userId, r); erro != nil {
		resposta.Erro(w, http.StatusForbidden, erro)
		return
	}

	erro = service.UpdatePublicacao(publicacao, userId)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	resposta.JSON(w, http.StatusNoContent, nil)

}

// Delete usuario
func DeletePublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}

	if erro = service.Permissoes(userId, r); erro != nil {
		resposta.Erro(w, http.StatusForbidden, erro)
		return
	}

	erro = service.DeletePublicacao(userId)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	resposta.JSON(w, http.StatusNoContent, nil)
}

// atualiza publicacao
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	erro = service.CurtirPublicacao(publicacaoId)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	resposta.JSON(w, http.StatusNoContent, nil)

}

// atualiza publicacao
func DesCurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	erro = service.DesCurtirPublicacao(publicacaoId)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	resposta.JSON(w, http.StatusNoContent, nil)

}

package controller

import (
	"api/src/autenticacao"
	"api/src/modelo"
	"api/src/resposta"
	"api/src/service"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Salvar usuario
func SaveUser(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	var usuario modelo.User
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
	}

	usuario, erro = service.Save(usuario)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	resposta.JSON(w, http.StatusOK, usuario)
}

// Buscar todos
func FindUsers(w http.ResponseWriter, r *http.Request) {
	userNameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))
	users, erro := service.FindUsersByNameOrNick(userNameOrNick)

	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
	}
	if users != nil {
		resposta.JSON(w, http.StatusOK, users)
	}
	resposta.JSON(w, http.StatusBadRequest, nil)
}

// Buscar por id
func FindUserPorId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	user, erro := service.FindUserById(int(userId))
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
	}
	resposta.JSON(w, http.StatusOK, user)
}

// Atualizar usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	userId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}

	var usuario modelo.User
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
	}

	if erro = service.Permissoes(userId, r); erro != nil {
		resposta.Erro(w, http.StatusForbidden, erro)
		return
	}

	erro = service.UpdateUserById(usuario, userId)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	resposta.JSON(w, http.StatusNoContent, nil)
}

// Delete usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}

	if erro = service.Permissoes(userId, r); erro != nil {
		resposta.Erro(w, http.StatusForbidden, erro)
		return
	}

	erro = service.DeleteUserById(userId)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	resposta.JSON(w, http.StatusNoContent, nil)
}

// Salvar usuario
func SeguirUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userAseguirId, erro := strconv.ParseUint(parametros["usuario_id"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}
	usuarioLogadoId, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if userAseguirId == usuarioLogadoId {
		resposta.Erro(w, http.StatusBadRequest, errors.New("Usuario não pode seguir se mesmo"))
		return
	}

	var seguidor modelo.Seguidor
	seguidor.UserID = usuarioLogadoId
	seguidor.SeguidorID = userAseguirId

	erro = service.SalvarSeguidor(seguidor)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}
	resposta.JSON(w, http.StatusNoContent, nil)
}

// deixar de seguir
func DeixarDeSeguir(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userAseguirId, erro := strconv.ParseUint(parametros["usuario_id"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioLogadoId, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if userAseguirId == usuarioLogadoId {
		resposta.Erro(w, http.StatusBadRequest, errors.New("Usuario não pode deixar de seguir se mesmo"))
		return
	}

	var seguidor modelo.Seguidor
	seguidor.UserID = usuarioLogadoId
	seguidor.SeguidorID = userAseguirId

	erro = service.DeixardeSeguir(seguidor)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		//}
		resposta.JSON(w, http.StatusNoContent, nil)
	}
}

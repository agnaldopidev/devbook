package rotas

import (
	"api/src/controller"
	"net/http"
)

var rotasLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controller.Login,
	RequerAutenticacao: false,
}

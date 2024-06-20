package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

// Rotas de login para buscar usuario na API
var loginRotas = Router{
	URI:                 "/fazer-login",
	Metodo:              http.MethodPost,
	Funcao:              controllers.FazerLogin,
	RequerAuthenticacao: false,
}

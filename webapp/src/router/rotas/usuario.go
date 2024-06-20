package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

// Rotas de login
var usuarioRotas = []Router{

	{
		URI:                 "/cadastro",
		Metodo:              http.MethodPost,
		Funcao:              controllers.Cadastrar,
		RequerAuthenticacao: false,
	},
}

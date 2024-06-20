package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

// Rotas de login
var publicacaoRotas = []Router{

	{
		URI:                 "/publicacoes",
		Metodo:              http.MethodPost,
		Funcao:              controllers.Publicar,
		RequerAuthenticacao: false,
	},
}

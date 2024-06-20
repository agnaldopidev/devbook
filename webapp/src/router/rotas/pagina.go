package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

// Rotas de login
var paginaRotas = []Router{
	{
		URI:                 "/",
		Metodo:              http.MethodGet,
		Funcao:              controllers.CarregaPagLogin,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/login",
		Metodo:              http.MethodGet,
		Funcao:              controllers.CarregaPagLogin,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/cadastro",
		Metodo:              http.MethodGet,
		Funcao:              controllers.CarregaPagCadastro,
		RequerAuthenticacao: false,
	},
	{
		URI:                 "/home",
		Metodo:              http.MethodGet,
		Funcao:              controllers.CarregaPagHome,
		RequerAuthenticacao: true,
	},
}

package rotas

import (
	"api/src/controller"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controller.SaveUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controller.FindUsers,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controller.FindUserPorId,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controller.UpdateUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeleteUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuario_id}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controller.SeguirUser,
		RequerAutenticacao: true,
	},
	{

		URI:                "/usuarios/{usuario_id}/deixar-de-seguir",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeixarDeSeguir,
		RequerAutenticacao: true,
	},
}

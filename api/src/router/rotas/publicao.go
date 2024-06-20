package rotas

import (
	"api/src/controller"
	"net/http"
)

var rotasPublicao = []Rota{
	{
		URI:                "/publicacao",
		Metodo:             http.MethodPost,
		Funcao:             controller.SalvarPublicacao,
		RequerAutenticacao: false,
	},
	{
		URI:                "/publicacao",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarPucacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarPublicacaoPorId,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controller.UpdatePublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeletePublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{id}/curtir",
		Metodo:             http.MethodPut,
		Funcao:             controller.DeletePublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacao/{id}/descurtir",
		Metodo:             http.MethodPut,
		Funcao:             controller.DeletePublicacao,
		RequerAutenticacao: true,
	},
}

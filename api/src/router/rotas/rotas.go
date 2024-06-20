package rotas

import (
	"api/src/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Estrutura que respresenta rotas
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configuracion router para
func Configure(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotasLogin)
	rotas = append(rotas, rotasPublicao...)

	for _, rota := range rotas {
		fmt.Println("Tested rotes", rota.URI, rota.Metodo)
		if rota.RequerAutenticacao {
			r.HandleFunc(
				rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}

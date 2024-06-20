package rotas

import (
	"fmt"
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Router struct {
	URI                 string
	Metodo              string
	Funcao              func(http.ResponseWriter, *http.Request)
	RequerAuthenticacao bool
}

// NewRouter returns a new
func ConfigureRouter(r *mux.Router) *mux.Router {
	rotas := paginaRotas
	rotas = append(rotas, loginRotas)
	rotas = append(rotas, usuarioRotas...)
	rotas = append(rotas, publicacaoRotas...)

	for _, rota := range rotas {
		fmt.Println("rotas: ", rota.URI, rota.Metodo)
		if rota.RequerAuthenticacao {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}

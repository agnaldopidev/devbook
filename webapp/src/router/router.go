package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

// Inicialization das rotas
func Init() *mux.Router {
	r := mux.NewRouter()
	return rotas.ConfigureRouter(r)
}

package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Criar as rotas da api
func Initialize() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configure(r)
}

package middlewares

import (
	"api/src/autenticacao"
	"api/src/resposta"
	"fmt"
	"net/http"
)

// Logger
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Autenticar user
func Autenticar(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			resposta.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}

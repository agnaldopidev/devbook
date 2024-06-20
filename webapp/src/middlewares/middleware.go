package middlewares

import (
	"fmt"
	"net/http"
	"webapp/src/cookie"
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
		if _, erro := cookie.LerValoresDoCookie(r); erro != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}
}

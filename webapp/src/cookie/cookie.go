package cookie

import (
	"net/http"
	"webapp/src/config"
	"webapp/src/modelo"

	"github.com/gorilla/securecookie"
)

var secure *securecookie.SecureCookie

// Inicializa: Inicialize secure cookies
func Inicializa() {
	secure = securecookie.New(config.Hashkey, config.Blockkey)
}

// SalvarNoCookie: Salva valores no cookie
func SalvarNoCookie(w http.ResponseWriter, dadoLogin modelo.Autenticacao) error {
	dados := map[string]string{
		"id":    dadoLogin.Id,
		"token": dadoLogin.Token,
	}
	dadosCodificados, erro := secure.Encode("dados", dados)
	if erro != nil {
		return erro
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

// LerValoresDoCookie: Retorna os valore armazenados no cookie
func LerValoresDoCookie(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}
	valores := make(map[string]string)
	if erro = secure.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}
	return valores, nil
}

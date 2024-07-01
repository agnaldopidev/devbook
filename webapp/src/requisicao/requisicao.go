package requisicao

import (
	"fmt"
	"io"
	"net/http"
	"webapp/src/cookie"
)

// FazerRequisicaoComAutenticacao faz uma requisicao com autorización
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}
	cookie, _ := cookie.LerValoresDoCookie(r)
	fmt.Println(cookie["token"])
	request.Header.Add("Authorization", "Bearer "+cookie["token"])
	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}
	return response, nil
}

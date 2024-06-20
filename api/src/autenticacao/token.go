package autenticacao

import (
	"api/src/config"
	"api/src/modelo"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(user modelo.User) (string, error) {
	// Defina a chave secreta usada para assinar o token
	secret := config.SecretKey
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Expira em 24 horas
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Assine o token com a chave secreta e obtenha a representação de string do token
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Erro ao assinar o token:", err)
		return "", err
	}

	//fmt.Println("Token JWT gerado:", tokenString)
	return tokenString, nil
}

// Validar token jwt das requisicao
func ValidarToken(r *http.Request) error {
	tokenStraindo := estrairToken(r)
	token, erro := jwt.Parse(tokenStraindo, retornaChaveSecretKey)
	if erro != nil {
		return erro
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Token invalido")
}

func estrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornaChaveSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de assinatura inesperado %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtrairUsuarioId(r *http.Request) (uint64, error) {
	tokenStraindo := estrairToken(r)
	token, erro := jwt.Parse(tokenStraindo, retornaChaveSecretKey)
	if erro != nil {
		return 0, erro
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ususrioId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return ususrioId, nil
	}
	return 0, errors.New("Token invalido")

}

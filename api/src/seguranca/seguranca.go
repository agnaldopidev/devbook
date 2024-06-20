package seguranca

import (
	"golang.org/x/crypto/bcrypt"
)

// Cria hash da senha
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// Compara hash da senha
func Compare(senhaHash string, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
}

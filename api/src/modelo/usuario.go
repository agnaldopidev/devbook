package modelo

import (
	"api/src/seguranca"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa a table user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Valida e limpa campos
func (user *User) Prepare(etapa string) error {
	if erro := user.validar(etapa); erro != nil {
		return erro
	}
	if erro := user.format(etapa); erro != nil {
		return erro
	}
	return nil
}

func (user *User) validar(etapa string) error {
	fmt.Println("nome", user.Nome)
	if user.Nome == "" {
		return errors.New("nome é obrigatorio")
	}

	if user.Email == "" {
		return errors.New("email é obrigatorio")
	}
	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("email é invalido")
	}
	if user.Nick == "" {
		return errors.New("nick é obrigatorio")
	}
	if etapa == "cadastro" {
		if user.Senha == "" {
			return errors.New("senha é obrigatorio")
		}
	}
	return nil
}

func (user *User) format(etapa string) error {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	user.Senha = strings.TrimSpace(user.Senha)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(user.Senha)
		if erro != nil {
			return erro
		}
		user.Senha = string(senhaComHash)
	}
	return nil
}

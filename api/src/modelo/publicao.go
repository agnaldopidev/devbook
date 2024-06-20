package modelo

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa a table publicacao
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorId,omitempty"`
	NickAutor string    `json:"nickAutor,omitempty"`
	Curtida   uint64    `json:"curtida,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Valida e limpa campos
func (publicacao *Publicacao) Prepare() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}
	publicacao.format()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("Titulo é obrigatorio")
	}

	if publicacao.Conteudo == "" {
		return errors.New("Conteúdo é obrigatorio")
	}
	return nil
}

func (publicacao *Publicacao) format() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}

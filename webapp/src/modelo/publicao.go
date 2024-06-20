package modelo

import (
	"time"
)

// Publicacao representa a table publicacao
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorId,omitempty"`
	NickAutor string    `json:"nickAutor,omitempty"`
	Curtida   uint64    `json:"curtida"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

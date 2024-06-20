package repository

import (
	"api/src/modelo"
	"database/sql"
)

// tipo usuario
type PublicacaoRepository struct {
	db *sql.DB
}

// novo usuario
func NewPublicacaoRepository(db *sql.DB) *PublicacaoRepository {
	return &PublicacaoRepository{db}
}

// salvar
func (resitorio *PublicacaoRepository) SalvarPublicacao(Publicacao modelo.Publicacao) (modelo.Publicacao, error) {
	query := `INSERT INTO PUBLICACAO (TITULO, CONTEUDO,AUTOR_ID) VALUES ($1, $2, $3) RETURNING id`
	stmt, erro := resitorio.db.Prepare(query)
	if erro != nil {
		return modelo.Publicacao{}, erro
	}
	defer stmt.Close()

	erro = stmt.QueryRow(Publicacao.Titulo, Publicacao.Conteudo, Publicacao.AutorId).Scan(&Publicacao.ID)
	if erro != nil {
		return modelo.Publicacao{}, erro
	}
	return Publicacao, nil
}

// busca Publicacao
func (repository PublicacaoRepository) BuscarPucacao(id uint64) ([]modelo.Publicacao, error) {
	query := `select
	distinct p.*,
	u.nick
	from
		publicacao p
	join usuario u on
		p.autor_id = u.id
	join seguidor s on
		p.autor_id = s.usuario_id
	where
	u.id = $1
	or s.seguidor_id = $2`
	linhas, erro := repository.db.Query(query, id, id)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var lista []modelo.Publicacao

	for linhas.Next() {
		var obj modelo.Publicacao
		if erro = linhas.Scan(&obj.ID,
			&obj.Titulo,
			&obj.Conteudo,
			&obj.AutorId,
			&obj.Curtida,
			&obj.CreatedAt,
			&obj.NickAutor,
		); erro != nil {
			return nil, erro
		}

		lista = append(lista, obj)

	}

	return lista, nil
}

// busca Publicacao por id
func (repository PublicacaoRepository) BuscarPublicacaoPorId(id uint64) ([]modelo.Publicacao, error) {
	query := `select
				p.id ,
				p.titulo,
				p.conteudo,
				p.curtida,
				p.createdat,
				(
				select
					u.nome
				from
					usuario u
				where
					u.id = p.autor_id) nick_autor
			from
				publicacao p
			where
				p.autor_id =$1`
	linhas, erro := repository.db.Query(query, id)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var lista []modelo.Publicacao

	for linhas.Next() {
		var obj modelo.Publicacao
		if erro = linhas.Scan(&obj.ID,
			&obj.Titulo,
			&obj.Conteudo,
			&obj.Curtida,
			&obj.CreatedAt,
			&obj.NickAutor,
		); erro != nil {
			return nil, erro
		}

		lista = append(lista, obj)

	}

	return lista, nil
}

// update user
func (repository PublicacaoRepository) UpdatePublicacao(publicacao modelo.Publicacao, ID uint64) error {
	query := `UPDATE publicacao SET titulo=$1,conteudo=$2,curtida=$3 WHERE ID = $4`
	stmt, erro := repository.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.Curtida, ID); erro != nil {
		return erro
	}
	return nil
}

// busca usuario por id
func (repository PublicacaoRepository) DeletePublicacao(ID uint64) error {
	query := `DELETE FROM publicacao WHERE ID = $1`
	stmt, erro := repository.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

// curtir publicacao
func (repository PublicacaoRepository) CurtirPublicacao(publicacaoID uint64) error {
	query := `UPDATE publicacao SET curtir=curtir+1 WHERE ID = $1`
	stmt, erro := repository.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

// curtir publicacao
func (repository PublicacaoRepository) DesCurtirPublicacao(publicacaoID uint64) error {
	query := `UPDATE publicacao SET case when curtir>0 then curtir=curtir-1 else 0 end WHERE ID = $1`
	stmt, erro := repository.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(publicacaoID); erro != nil {
		return erro
	}
	return nil
}

package service

import (
	"api/src/database"
	"api/src/modelo"
	"api/src/repository"
)

// salvar publicacao
func SalvarPublicacao(publicacao modelo.Publicacao) (modelo.Publicacao, error) {
	if erro := publicacao.Prepare(); erro != nil {
		return modelo.Publicacao{}, erro
	}
	db, erro := database.InitializePostgres()
	if erro != nil {
		return modelo.Publicacao{}, erro
	}
	defer db.Close()
	pubRepository := repository.NewPublicacaoRepository(db)
	publicacao, erro = pubRepository.SalvarPublicacao(publicacao)
	if erro != nil {
		return modelo.Publicacao{}, erro
	}
	return publicacao, nil
}

// buscar publicacao
func BuscarPucacao(id uint64) ([]modelo.Publicacao, error) {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()
	repo := repository.NewPublicacaoRepository(db)
	lista, erro := repo.BuscarPucacao(id)
	if erro != nil {
		return nil, erro
	}

	return lista, nil
}

// buscar publicacao por id
func BuscarPublicacaoPorId(id uint64) ([]modelo.Publicacao, error) {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()
	repo := repository.NewPublicacaoRepository(db)
	lista, erro := repo.BuscarPublicacaoPorId(id)
	if erro != nil {
		return nil, erro
	}

	return lista, nil
}

// update Publicacao
func UpdatePublicacao(publicacao modelo.Publicacao, ID uint64) error {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return erro
	}

	defer db.Close()
	repository := repository.NewPublicacaoRepository(db)
	erro = repository.UpdatePublicacao(publicacao, ID)
	if erro != nil {
		return erro
	}
	return nil
}

// delete
func DeletePublicacao(ID uint64) error {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return erro
	}

	defer db.Close()
	repository := repository.NewPublicacaoRepository(db)
	erro = repository.DeletePublicacao(ID)
	if erro != nil {
		return erro
	}
	return nil
}

// cursit publicacao
func CurtirPublicacao(publicacaoID uint64) error {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return erro
	}

	defer db.Close()
	repository := repository.NewPublicacaoRepository(db)
	erro = repository.CurtirPublicacao(publicacaoID)
	if erro != nil {
		return erro
	}
	return nil
}

// descurtir publicacao
func DesCurtirPublicacao(publicacaoID uint64) error {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return erro
	}

	defer db.Close()
	repository := repository.NewPublicacaoRepository(db)
	erro = repository.DesCurtirPublicacao(publicacaoID)
	if erro != nil {
		return erro
	}
	return nil
}

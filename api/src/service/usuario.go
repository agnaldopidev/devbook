package service

import (
	"api/src/autenticacao"
	"api/src/database"
	"api/src/modelo"
	"api/src/repository"
	"errors"
	"net/http"
)

// salvar user
func Save(user modelo.User) (modelo.User, error) {
	if erro := user.Prepare("cadastro"); erro != nil {
		return modelo.User{}, erro
	}

	db, erro := database.InitializePostgres()
	if erro != nil {
		return modelo.User{}, erro
	}
	defer db.Close()
	userRepository := repository.NewUsuarioRepository(db)
	user, erro = userRepository.Save(user)
	if erro != nil {
		return modelo.User{}, erro
	}
	return user, nil
}

// buscar usuario
func FindUsersByNameOrNick(userNameOrNick string) ([]modelo.User, error) {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()
	userRepository := repository.NewUsuarioRepository(db)
	users, erro := userRepository.FindUsersByNameOrNick(userNameOrNick)
	if erro != nil {
		return nil, erro
	}

	return users, nil
}

// buscar po id
func FindUserById(ID int) (modelo.User, error) {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return modelo.User{}, erro
	}

	defer db.Close()
	repository := repository.NewUsuarioRepository(db)
	user, erro := repository.FindUserById(ID)
	if erro != nil {
		return modelo.User{}, erro
	}
	return user, erro
}

// buscar
func UpdateUserById(userUp modelo.User, ID uint64) error {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return erro
	}

	defer db.Close()
	repository := repository.NewUsuarioRepository(db)
	erro = repository.UpdateUserById(userUp, ID)
	if erro != nil {
		return erro
	}
	return nil
}

// delete
func DeleteUserById(ID uint64) error {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return erro
	}

	defer db.Close()
	repository := repository.NewUsuarioRepository(db)
	erro = repository.DeleteUserById(ID)
	if erro != nil {
		return erro
	}
	return nil
}

// buscar po email
func FindUserByEmail(email string) (modelo.User, error) {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return modelo.User{}, erro
	}

	defer db.Close()
	repository := repository.NewUsuarioRepository(db)
	userDb, erro := repository.FindUserByEmail(email)
	if erro != nil {
		return modelo.User{}, erro
	}
	return userDb, erro
}

// Verifica permissao do usuario
func Permissoes(userID uint64, r *http.Request) error {
	// buscar id do usuario logado
	usuarioLogadoId, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		return erro
	}

	// usuario logado nao pode atualiar informacao de outro usuario
	if usuarioLogadoId != userID {
		return errors.New("Permissão negada para usuario")
	}
	return nil
}

// salvar seguidor
func SalvarSeguidor(seguidor modelo.Seguidor) error {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return erro
	}
	defer db.Close()
	seguidorRepository := repository.NewUsuarioRepository(db)
	erro = seguidorRepository.SalvarSeguidor(seguidor)
	if erro != nil {
		return erro
	}
	return nil
}

// deixae de seguir
func DeixardeSeguir(seguidor modelo.Seguidor) error {
	db, erro := database.InitializePostgres()
	if erro != nil {
		return erro
	}
	defer db.Close()
	seguidorRepository := repository.NewUsuarioRepository(db)
	erro = seguidorRepository.DeixardeSeguir(seguidor)
	if erro != nil {
		return erro
	}
	return nil
}

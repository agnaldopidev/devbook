package repository

import (
	"api/src/modelo"
	"database/sql"
	"fmt"
)

// tipo usuario
type UserRepository struct {
	db *sql.DB
}

// novo usuario
func NewUsuarioRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// salvar
func (resitorio *UserRepository) Save(Usuario modelo.User) (modelo.User, error) {
	query := `INSERT INTO usuario (nome, nick,email,senha) VALUES ($1, $2, $3, $4) RETURNING id`
	stmt, erro := resitorio.db.Prepare(query)
	if erro != nil {
		return modelo.User{}, erro
	}
	defer stmt.Close()

	erro = stmt.QueryRow(Usuario.Nome, Usuario.Nick, Usuario.Email, Usuario.Senha).Scan(&Usuario.ID)
	if erro != nil {
		return modelo.User{}, erro
	}
	return Usuario, nil
}

// busca usuario
func (repository UserRepository) FindUsersByNameOrNick(userNameOrNick string) ([]modelo.User, error) {
	query := `SELECT id,nome,email,nick,createdat FROM usuario WHERE nome LIKE $1 OR nick LIKE $2`
	userNameOrNick = fmt.Sprintf("%%%s%%", userNameOrNick) //%userNameOrNick%
	linhas, erro := repository.db.Query(query, userNameOrNick, userNameOrNick)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var users []modelo.User

	for linhas.Next() {
		var user modelo.User
		if erro = linhas.Scan(&user.ID,
			&user.Nome,
			&user.Email,
			&user.Nick,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)

	}

	return users, nil
}

// busca usuario por id
func (repository UserRepository) FindUserById(ID int) (modelo.User, error) {
	query := `select id,nome,nick,email,createdat from usuario where id=$1`
	linhas, erro := repository.db.Query(query, ID)
	if erro != nil {
		return modelo.User{}, erro

	}
	defer linhas.Close()

	var user modelo.User

	for linhas.Next() {
		if erro = linhas.Scan(&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return modelo.User{}, erro
		}

	}

	return user, nil
}

// update user
func (repository UserRepository) UpdateUserById(userUp modelo.User, ID uint64) error {
	query := `UPDATE usuario SET nome = $1, email = $2, nick = $3 WHERE ID = $4`
	stmt, erro := repository.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(userUp.Nome, userUp.Email, userUp.Nick, ID); erro != nil {
		return erro
	}
	return nil
}

// busca usuario por id
func (repository UserRepository) DeleteUserById(ID uint64) error {
	query := `DELETE FROM usuario WHERE ID = $1`
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

// busca usuario por email
func (repository UserRepository) FindUserByEmail(email string) (modelo.User, error) {
	fmt.Println("\nemail", email)
	query := `select id,email,senha from usuario where email=$1`
	linhas, erro := repository.db.Query(query, email)
	if erro != nil {
		return modelo.User{}, erro

	}
	defer linhas.Close()

	var user modelo.User

	for linhas.Next() {
		if erro = linhas.Scan(&user.ID,
			&user.Email,
			&user.Senha,
		); erro != nil {
			return modelo.User{}, erro
		}

	}

	return user, nil
}

// salvar
func (resitorio *UserRepository) SalvarSeguidor(Seguidor modelo.Seguidor) error {
	query := `INSERT INTO seguidor (usuario_id, seguidor_id) 
	VALUES ($1, $2) ON CONFLICT (usuario_id,seguidor_id) DO NOTHING`
	stmt, erro := resitorio.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	_, erro = stmt.Exec(Seguidor.UserID, Seguidor.SeguidorID)
	if erro != nil {
		return erro
	}
	return nil
}

// deixar de seguir
func (resitorio *UserRepository) DeixardeSeguir(Seguidor modelo.Seguidor) error {
	query := `DELETE FROM seguidor WHERE usuario_id=$1 and seguidor_id=$2`
	stmt, erro := resitorio.db.Prepare(query)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	_, erro = stmt.Exec(Seguidor.UserID, Seguidor.SeguidorID)
	if erro != nil {
		return erro
	}
	return nil
}

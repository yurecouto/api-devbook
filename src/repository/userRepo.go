package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repo Users) Create(usuario models.User) (uint64, error) {
	statement, erro := repo.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(usuario.Name, usuario.Nick, usuario.Email, usuario.Password)
	if erro != nil {
		return 0, erro
	}

	lastId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastId), nil
}

func (repo Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, erro := repo.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo Users) SearchByID(ID uint64) (models.User, error) {
	lines, erro := repo.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		ID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repo Users) Update(ID uint64, user models.User) error {
	statement, erro := repo.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

func (repo Users) Delete(ID uint64) error {
	statement, erro := repo.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

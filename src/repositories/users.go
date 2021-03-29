package repositories

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

func (repository Users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare(
		"insert into user (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(
		user.Name,
		user.Nick,
		user.Email,
		user.Password,
	)
	if error != nil {
		return 0, error
	}

	lastInsertedID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedID), nil
}

func (repository Users) FindAll(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	users, error := repository.db.Query(
		"select id, name, nick, email, createdAt from user where name LIKE ? OR nick LIKE ?",
		nameOrNick,
		nameOrNick,
	)
	if error != nil {
		return nil, error
	}
	defer users.Close()

	var findUsers []models.User

	for users.Next() {
		var user models.User

		if error = users.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}

		findUsers = append(findUsers, user)
	}

	return findUsers, nil
}

func (repository Users) FindOneById(userId uint64) (models.User, error) {
	user, error := repository.db.Query(
		"select id, name, nick, email, createdAt from user where id = ?",
		userId,
	)
	if error != nil {
		return models.User{}, error
	}
	defer user.Close()

	var findUser models.User

	if user.Next() {
		if error = user.Scan(
			&findUser.ID,
			&findUser.Name,
			&findUser.Nick,
			&findUser.Email,
			&findUser.CreatedAt,
		); error != nil {
			return models.User{}, error
		}
	}

	return findUser, nil
}

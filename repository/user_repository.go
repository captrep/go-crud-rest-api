package repository

import (
	"database/sql"
	"log"
	"strings"

	"github.com/captrep/go-crud-rest-api/model/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

const (
	queryInsertUser = "INSERT INTO users (id, first_name, last_name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	queryFindAll    = "SELECT * FROM users"
	queryGetById    = "SELECT id, first_name, last_name, email, created_at, updated_at FROM users WHERE id=?"
	queryUpdate     = "UPDATE users SET first_name = ?, last_name = ?, email = ?, updated_at = ? WHERE id=?"
	queryDelete     = "DELETE FROM users WHERE id=?"
	errorNoRows     = "no rows in result set"
)

func (repository *UserRepository) Save(user domain.User) (domain.User, error) {
	statement, err := repository.db.Prepare(queryInsertUser)
	if err != nil {
		return user, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Id, user.Firstname, user.Lastname, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return user, err
	}
	log.Println(result)

	return user, nil
}

func (repository *UserRepository) GetAll() ([]domain.User, error) {
	statement, err := repository.db.Prepare(queryFindAll)
	if err != nil {
		return nil, err
	}

	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]domain.User, 0)
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository *UserRepository) FindById(userId string) (domain.User, error) {
	statement, err := repository.db.Prepare(queryGetById)
	user := domain.User{}
	if err != nil {
		return user, err
	}
	defer statement.Close()

	result := statement.QueryRow(userId)
	if err := result.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return user, err
		}
		return user, err
	}

	return user, nil

}

func (repository *UserRepository) Update(user domain.User) (domain.User, error) {
	statement, err := repository.db.Prepare(queryUpdate)
	if err != nil {
		return user, err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Firstname, user.Lastname, user.Email, user.UpdatedAt, user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *UserRepository) Delete(user domain.User) error {
	statement, err := repository.db.Prepare(queryDelete)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Id)
	if err != nil {
		return err
	}

	return nil
}

package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (int64, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, username, password) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Email, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return int64(id), nil
}

func (r *AuthPostgres) GetUser(email, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}

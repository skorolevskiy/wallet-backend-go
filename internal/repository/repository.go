package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"github.com/skorolevskiy/wallet-backend-go/internal/repository/postgres"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Wallet interface {
}

type Transaction interface {
}

type Repository struct {
	Authorization
	Wallet
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
	}
}

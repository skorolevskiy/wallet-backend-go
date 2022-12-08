package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}

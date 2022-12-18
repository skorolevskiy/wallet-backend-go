package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"github.com/skorolevskiy/wallet-backend-go/internal/repository/postgres"
)

type Authorization interface {
	CreateUser(user domain.User) (int64, error)
	GetUser(username, password string) (domain.User, error)
}

type Tokens interface {
	Create(token domain.RefreshToken) error
	Get(token string) (domain.RefreshToken, error)
}

type Wallet interface {
	CreateWallet(userId int64, wallet domain.Wallet) (int64, error)
	GetAllWallets(userId int64) ([]domain.Wallet, error)
	GetWalletById(userId, walletId int64) (domain.Wallet, error)
	UpdateWallet(userId, walletId int64, input domain.UpdateWalletInput) error
	DeleteWallet(userId, walletId int64) error
}

type Transaction interface {
	CreateTransaction(walletId, userId int64, transaction domain.Transaction) (int64, error)
	GetAllTransactions(walletId int64) ([]domain.Transaction, error)
	GetTransactionById(userId, transactionId int64) (domain.Transaction, error)
	UpdateTransaction(userId, walletId, transactionId int64, input domain.UpdateTransactionInput, amount float64) error
	UpdateTransactionBalance(walletId, transactionId int64, amount float64) error
	DeleteTransaction(userId, walletId, transactionId int64, amount float64) error
}

type Repository struct {
	Authorization
	Tokens
	Wallet
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		Tokens:        postgres.NewTokens(db),
		Wallet:        postgres.NewWalletPostgres(db),
		Transaction:   postgres.NewTransactionPostgres(db),
	}
}

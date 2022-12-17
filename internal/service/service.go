package service

import (
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"github.com/skorolevskiy/wallet-backend-go/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int64, error)
	SignIn(username, password string) (string, string, error)
	ParseToken(token string) (int64, error)
	RefreshToken(refreshToken string) (string, string, error)
}

type Wallet interface {
	CreateWallet(userId int64, wallet domain.Wallet) (int64, error)
	GetAllWallets(userId int64) ([]domain.Wallet, error)
	GetWalletById(userId, walletId int64) (domain.Wallet, error)
	UpdateWallet(userId, walletId int64, input domain.UpdateWalletInput) error
	DeleteWallet(userId, walletId int64) error
}

type Transaction interface {
	CreateTransaction(userId, walletId int64, transaction domain.Transaction) (int64, error)
	GetAllTransactions(userId, walletId int64) ([]domain.Transaction, error)
	GetTransactionById(userId, walletId, transactionId int64) (domain.Transaction, error)
}

type Service struct {
	Authorization
	Wallet
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, repos.Tokens),
		Wallet:        NewWalletService(repos.Wallet),
		Transaction:   NewTransactionService(repos.Transaction, repos.Wallet),
	}
}

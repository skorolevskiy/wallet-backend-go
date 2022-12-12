package service

import (
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"github.com/skorolevskiy/wallet-backend-go/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Wallet interface {
}

type Transaction interface {
}

type Service struct {
	Authorization
	Wallet
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

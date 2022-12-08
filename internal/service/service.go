package service

import "github.com/skorolevskiy/wallet-backend/internal/repository"

type Authorization interface {
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
	return &Service{}
}

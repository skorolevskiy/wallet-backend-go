package service

import (
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"github.com/skorolevskiy/wallet-backend-go/internal/repository"
)

type WalletService struct {
	repo repository.Wallet
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) CreateWallet(userId int64, wallet domain.Wallet) (int64, error) {
	wallet.Balance = 0

	return s.repo.CreateWallet(userId, wallet)
}

func (s *WalletService) GetAllWallets(userId int64) ([]domain.Wallet, error) {
	return s.repo.GetAllWallets(userId)
}

func (s *WalletService) GetWalletById(userId, walletId int64) (domain.Wallet, error) {
	return s.repo.GetWalletById(userId, walletId)
}

func (s *WalletService) UpdateWallet(userId, walletId int64, input domain.UpdateWalletInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateWallet(userId, walletId, input)
}

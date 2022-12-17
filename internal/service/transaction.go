package service

import (
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"github.com/skorolevskiy/wallet-backend-go/internal/repository"
)

type TransactionService struct {
	repo       repository.Transaction
	walletRepo repository.Wallet
}

func NewTransactionService(repo repository.Transaction, walletRepo repository.Wallet) *TransactionService {
	return &TransactionService{repo: repo, walletRepo: walletRepo}
}

func (s *TransactionService) CreateTransaction(userId, walletId int64, transaction domain.Transaction) (int64, error) {
	wallet, err := s.walletRepo.GetWalletById(userId, walletId)
	if err != nil {
		return 0, err
	}
	walletAmount := wallet.Balance
	transaction.BalanceAfter = walletAmount + transaction.Amount - transaction.CommissionAmount
	return s.repo.CreateTransaction(walletId, userId, transaction)
}

func (s *TransactionService) GetAllTransactions(userId, walletId int64) ([]domain.Transaction, error) {
	_, err := s.walletRepo.GetWalletById(userId, walletId)
	if err != nil {
		return nil, err
	}
	return s.repo.GetAllTransactions(walletId)
}

func (s *TransactionService) GetTransactionById(userId, walletId, transactionId int64) (domain.Transaction, error) {
	_, err := s.walletRepo.GetWalletById(userId, walletId)
	if err != nil {
		return domain.Transaction{}, err
	}
	return s.repo.GetTransactionById(walletId, transactionId)
}

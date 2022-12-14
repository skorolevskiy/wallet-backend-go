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

func (s *TransactionService) UpdateTransaction(userId, walletId, transactionId int64, input domain.UpdateTransactionInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	_, err := s.walletRepo.GetWalletById(userId, walletId)
	if err != nil {
		return err
	}
	var newTransactionChangeAmount float64
	if input.Amount != nil {
		transactionDB, err := s.repo.GetTransactionById(walletId, transactionId)
		if err != nil {
			return err
		}
		if input.Amount != &transactionDB.Amount && input.CommissionAmount != &transactionDB.CommissionAmount {
			if input.CommissionAmount != &transactionDB.CommissionAmount {
				newCommissionChangeAmount := *input.CommissionAmount - transactionDB.CommissionAmount
				newTransactionChangeAmount = *input.Amount - transactionDB.Amount - newCommissionChangeAmount
				newBalance := transactionDB.BalanceAfter + newTransactionChangeAmount
				input.BalanceAfter = &newBalance
			} else {
				newTransactionChangeAmount = *input.Amount - transactionDB.Amount
				newBalance := transactionDB.BalanceAfter + newTransactionChangeAmount
				input.BalanceAfter = &newBalance
			}
			err = s.repo.UpdateTransactionBalance(walletId, transactionId, newTransactionChangeAmount)
			if err != nil {
				return err
			}
		}
	}

	return s.repo.UpdateTransaction(userId, walletId, transactionId, input, newTransactionChangeAmount)
}

func (s *TransactionService) DeleteTransaction(userId, walletId, transactionId int64) error {
	_, err := s.walletRepo.GetWalletById(userId, walletId)
	if err != nil {
		return err
	}
	transactionDB, err := s.repo.GetTransactionById(walletId, transactionId)
	if err != nil {
		return err
	}
	amount := transactionDB.Amount - transactionDB.CommissionAmount
	err = s.repo.UpdateTransactionBalance(walletId, transactionId, -amount)
	if err != nil {
		return err
	}
	return s.repo.DeleteTransaction(userId, walletId, transactionId, amount)
}

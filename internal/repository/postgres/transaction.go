package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) CreateTransaction(walletId, userId int64, transaction domain.Transaction) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int64
	createTransactionQuery := fmt.Sprintf("INSERT INTO %s (wallet_id, description, amount, balance_after, commission_amount, currency) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", transactionTable)
	row := tx.QueryRow(createTransactionQuery, walletId, transaction.Description, transaction.Amount, transaction.BalanceAfter, transaction.CommissionAmount, transaction.Currency)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createWalletUpdateQuery := fmt.Sprintf("UPDATE %s SET balance=$1 WHERE id=$2 AND user_id=$3", walletsTable)
	_, err = r.db.Exec(createWalletUpdateQuery, transaction.BalanceAfter, walletId, userId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

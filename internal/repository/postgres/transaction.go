package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"strings"
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

	walletUpdateQuery := fmt.Sprintf("UPDATE %s SET balance=$1 WHERE id=$2 AND user_id=$3", walletsTable)
	_, err = tx.Exec(walletUpdateQuery, transaction.BalanceAfter, walletId, userId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TransactionPostgres) GetAllTransactions(walletId int64) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	getAllQuery := fmt.Sprintf("SELECT id, wallet_id, description, amount, balance_after, commission_amount, currency, created_at FROM %s WHERE wallet_id = $1", transactionTable)
	err := r.db.Select(&transactions, getAllQuery, walletId)
	return transactions, err
}

func (r *TransactionPostgres) GetTransactionById(walletId, transactionId int64) (domain.Transaction, error) {
	var transaction domain.Transaction
	getByIdQuery := fmt.Sprintf("SELECT id, wallet_id, description, amount, balance_after, commission_amount, currency, created_at FROM %s WHERE wallet_id = $1 AND id = $2", transactionTable)
	err := r.db.Get(&transaction, getByIdQuery, walletId, transactionId)
	return transaction, err
}

func (r *TransactionPostgres) UpdateTransaction(userId, walletId, transactionId int64, input domain.UpdateTransactionInput, amount float64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("amount=$%d", argId))
		args = append(args, *input.Amount)
		argId++
	}
	if input.BalanceAfter != nil {
		setValues = append(setValues, fmt.Sprintf("balance_after=$%d", argId))
		args = append(args, *input.BalanceAfter)
		argId++
	}
	if input.CommissionAmount != nil {
		setValues = append(setValues, fmt.Sprintf("commission_amount=$%d", argId))
		args = append(args, *input.CommissionAmount)
		argId++
	}
	if input.Currency != nil {
		setValues = append(setValues, fmt.Sprintf("currency=$%d", argId))
		args = append(args, *input.Currency)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	updateQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND wallet_id=$%d", transactionTable, setQuery, argId, argId+1)
	args = append(args, transactionId, walletId)
	logrus.Debugf("update Query: %s", updateQuery)
	logrus.Debugf("args: %s", args)
	_, err = tx.Exec(updateQuery, args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	createWalletUpdateQuery := fmt.Sprintf("UPDATE %s SET balance=balance+$1 WHERE id=$2 AND user_id=$3", walletsTable)
	_, err = tx.Exec(createWalletUpdateQuery, amount, walletId, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *TransactionPostgres) UpdateTransactionBalance(walletId, transactionId int64, amount float64) error {
	getHigherIdQuery := fmt.Sprintf("SELECT id, balance_after FROM %s WHERE wallet_id = $1 AND id > $2", transactionTable)
	rows, err := r.db.Query(getHigherIdQuery, walletId, transactionId)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var transactionIdChange int64
		var balanceAfterOld float64
		if err := rows.Scan(&transactionIdChange, &balanceAfterOld); err != nil {
			return err
		}
		updateAmountQuery := fmt.Sprintf("UPDATE %s SET balance_after = $1 WHERE id = $2", transactionTable)
		_, err := r.db.Exec(updateAmountQuery, balanceAfterOld+amount, transactionIdChange)
		if err != nil {
			return err
		}
	}

	return err
}

func (r *TransactionPostgres) DeleteTransaction(userId, walletId, transactionId int64, amount float64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE wallet_id=$1 AND id=$2", transactionTable)
	_, err = tx.Exec(deleteQuery, walletId, transactionId)
	if err != nil {
		tx.Rollback()
		return err
	}
	createWalletUpdateQuery := fmt.Sprintf("UPDATE %s SET balance=balance-$1 WHERE id=$2 AND user_id=$3", walletsTable)
	_, err = tx.Exec(createWalletUpdateQuery, amount, walletId, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

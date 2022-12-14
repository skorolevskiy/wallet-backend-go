package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
	"strings"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func NewWalletPostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}

func (r *WalletPostgres) CreateWallet(userId int64, wallet domain.Wallet) (int64, error) {
	var id int64
	createWalletQuery := fmt.Sprintf("INSERT INTO %s (user_id, name, balance, currency) VALUES ($1, $2, $3, $4) RETURNING id", walletsTable)
	err := r.db.QueryRow(createWalletQuery, userId, wallet.Name, wallet.Balance, wallet.Currency).Scan(&id)

	return id, err
}

func (r *WalletPostgres) GetAllWallets(userId int64) ([]domain.Wallet, error) {
	var wallets []domain.Wallet
	getAllQuery := fmt.Sprintf("SELECT id, user_id, name, balance, currency, register_at FROM %s WHERE user_id = $1", walletsTable)
	err := r.db.Select(&wallets, getAllQuery, userId)
	return wallets, err
}

func (r *WalletPostgres) GetWalletById(userId, walletId int64) (domain.Wallet, error) {
	var wallet domain.Wallet
	getByIdQuery := fmt.Sprintf("SELECT id, user_id, name, balance, currency, register_at FROM %s WHERE user_id = $1 AND id = $2", walletsTable)
	err := r.db.Get(&wallet, getByIdQuery, userId, walletId)
	return wallet, err
}

func (r *WalletPostgres) UpdateWallet(userId, walletId int64, input domain.UpdateWalletInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Currency != nil {
		setValues = append(setValues, fmt.Sprintf("currency=$%d", argId))
		args = append(args, *input.Currency)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	updateQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND user_id=$%d", walletsTable, setQuery, argId, argId+1)

	args = append(args, walletId, userId)
	logrus.Debug("update Query: %s", updateQuery)
	logrus.Debug("args: %s", args)
	_, err := r.db.Exec(updateQuery, input.Name, input.Currency, walletId, userId)
	return err
}

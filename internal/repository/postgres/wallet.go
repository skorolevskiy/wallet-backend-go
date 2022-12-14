package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
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

package domain

import "time"

type Transaction struct {
	Id               int64     `json:"id" db:"id"`
	WalletID         int64     `json:"walletId" db:"wallet_id"`
	Description      string    `json:"description" db:"description"`
	Amount           float64   `json:"amount" db:"amount"`
	BalanceAfter     float64   `json:"balanceAfter" db:"balance_after"`
	CommissionAmount float64   `json:"commissionAmount" db:"commission_amount"`
	Currency         string    `json:"currency" db:"currency"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
}

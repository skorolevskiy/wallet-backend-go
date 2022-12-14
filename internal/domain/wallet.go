package domain

import (
	"errors"
	"time"
)

type Wallet struct {
	ID         int       `json:"id" db:"id"`
	UserId     int64     `json:"user_id" db:"user_id"`
	Name       string    `json:"name" db:"name" binding:"required"`
	Balance    float64   `json:"balance" db:"balance"`
	Currency   string    `json:"currency" db:"currency"`
	RegisterAt time.Time `json:"register_at" db:"register_at"`
}

type UpdateWalletInput struct {
	Name     *string `json:"name"`
	Currency *string `json:"currency"`
}

func (i UpdateWalletInput) Validate() error {
	if i.Name == nil && i.Currency == nil {
		return errors.New("update structure not value")
	}

	return nil
}

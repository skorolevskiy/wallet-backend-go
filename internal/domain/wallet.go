package domain

import "time"

type Wallet struct {
	ID         int       `json:"id" db:"id"`
	UserId     int       `json:"user_id" db:"userId"`
	Name       string    `json:"name" db:"name" binding:"required"`
	Balance    float64   `json:"balance" db:"balance"`
	Currency   string    `json:"currency" db:"currency"`
	RegisterAt time.Time `json:"register_at" db:"registerAt"`
}

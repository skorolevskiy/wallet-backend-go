package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/skorolevskiy/wallet-backend-go/internal/domain"
)

type Tokens struct {
	db *sqlx.DB
}

func NewTokens(db *sqlx.DB) *Tokens {
	return &Tokens{db: db}
}

func (r *Tokens) Create(token domain.RefreshToken) error {
	_, err := r.db.Exec("INSERT INTO refresh_tokens (user_id, token, expires_at) values ($1, $2, $3)",
		token.UserID, token.Token, token.ExpiresAt)

	return err
}

func (r *Tokens) Get(token string) (domain.RefreshToken, error) {
	var t domain.RefreshToken
	err := r.db.QueryRow("SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token=$1", token).
		Scan(&t.ID, &t.UserID, &t.Token, &t.ExpiresAt)
	if err != nil {
		return t, err
	}
	_, err = r.db.Exec("DELETE FROM refresh_tokens WHERE user_id=$1", t.UserID)

	return t, err
}

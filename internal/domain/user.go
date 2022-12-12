package domain

import "time"

type User struct {
	ID         int64  `json:"-" db:"id"`
	Email      string `json:"email" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RegisterAt time.Time
}

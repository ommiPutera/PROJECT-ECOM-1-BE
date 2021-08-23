package model

import (
	"time"
)

type Transaction struct {
	ID            int       `json:"id"`
	FromEWalletID int       `json:"from_ewallet_id"`
	ToEWalletID   int       `json:"to_ewallet_id"`
	TotalAmount   int       `json:"total_amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IsDeleted     bool      `json:"is_deleted"`
}

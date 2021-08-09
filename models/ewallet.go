package model

import "time"

type EWallet struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

package model

import "time"

type Chat struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ShopID    string    `json:"shop_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

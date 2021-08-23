package model

import "time"

type Address struct {
	ID         int       `json:"id"`
	UserID     string    `json:"user_id"`
	ShopID     string    `json:"shop_id"`
	Location   string    `json:"location"`
	City       string    `json:"city"`
	PostalCode string    `json:"postal_code"`
	Lattitude  string    `json:"lattitude"`
	Longitude  string    `json:"longitude"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsDefault  bool      `json:"is_default"`
	IsDeleted  bool      `json:"is_deleted"`
}

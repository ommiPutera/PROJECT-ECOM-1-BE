package model

import "time"

type Order struct {
	ID             int       `json:"id"`
	ShopID         string    `json:"shop_id"`
	CategoryID     int       `json:"category_id"`
	TransactionID  int       `json:"amount_sold"`
	ProductImageID int       `json:"product_image_id"`
	SpecID         int       `json:"spec_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Price          uint      `json:"price"`
	Discount       uint      `json:"discount"`
	IsCashback     bool      `json:"is_cashback"`
	Quantity       uint      `json:"stock"`
	Weight         uint      `json:"weight"`
	ShippingFee    uint      `json:"shipping_fee"`
	OrderStatus    uint      `json:"order_status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	IsDeleted      bool      `json:"is_deleted"`
}

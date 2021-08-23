package model

type Cart struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	UserID    string `json:"user_id"`
	ShopID    string `json:"shop_id"`
}

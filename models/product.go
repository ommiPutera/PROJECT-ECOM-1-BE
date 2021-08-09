package model

import "time"

type Product struct {
	ID           string    `json:"id" bson:"_id"`
	ShopID       string    `json:"shop_id" bson:"shop_id"`
	CategoryID   string    `json:"category_id" bson:"category_id"`
	CategoryName string    `json:"category_name"`
	SpecID       string    `json:"spec_id" bson:"spec_id"`
	Name         string    `json:"name" bson:"name"`
	Description  string    `json:"description" bson:"description"`
	Price        uint      `json:"price" bson:"price"`
	Stock        uint      `json:"stock" bson:"stock"`
	Weight       uint      `json:"weight" bson:"weight"`
	AmountSold   uint      `json:"amount_sold" bson:"amount_sold"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
	IsDeleted    bool      `json:"is_deleted" bson:"is_deleted"`
}

type ProductCategory struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    ` json:"category" bson:"category"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type ProductPromotion struct {
	ID          int       `json:"id"`
	ProductID   int       `json:"product_id"`
	CategoryID  int       `json:"category_id"`
	Discount    uint      `json:"discount"`
	MinPurchase uint      `json:"min_purchase"`
	IsCashback  bool      `json:"is_cashback"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiredAt   time.Time `json:"expired_at"`
	IsDeleted   bool      `json:"is_deleted"`
}

package cart

import (
	"context"
	model "ecomjc-be/models"
)

type CartService interface {
	CreateCart(ctx context.Context, Cart *model.Cart) (*model.Cart, error)
	GetCarts(ctx context.Context) ([]*model.Cart, error)
	GetCartByID(ctx context.Context, CartID int) (*model.Cart, error)
	UpdateCart(ctx context.Context, Cart *model.Cart) (*model.Cart, error)
	DeleteCart(ctx context.Context, CartID int) (bool, error)
}

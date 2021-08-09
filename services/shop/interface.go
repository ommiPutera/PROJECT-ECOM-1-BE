package shop

import (
	"context"
	model "ecomjc-be/models"
)

type ShopService interface {
	CreateShop(ctx context.Context, shop *model.Shop) (*model.Shop, error)
	GetShops(ctx context.Context) ([]*model.Shop, error)
	GetShopByID(ctx context.Context, shopID string) ([]*model.Shop, error)
	UpdateShop(ctx context.Context, shop *model.Shop) (*model.Shop, error)
	DeleteShop(ctx context.Context, shopID string) (bool, error)
}

package order

import (
	"context"
	model "ecomjc-be/models"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error)
	GetOrders(ctx context.Context) ([]*model.Order, error)
	GetOrderByID(ctx context.Context, orderID int) (*model.Order, error)
	UpdateOrder(ctx context.Context, order *model.Order) (*model.Order, error)
	DeleteOrder(ctx context.Context, orderID int) (bool, error)
}

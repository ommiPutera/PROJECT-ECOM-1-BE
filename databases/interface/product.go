package database

import (
	"context"
	model "ecomjc-be/models"
)

type Product interface {
	CreateProduct(ctx context.Context, product *model.Product) (string, error)
	GetProducts(ctx context.Context) ([]*model.Product, error)
	GetProductByID(ctx context.Context, productID string) (*model.Product, error)

	CreateProductCategory(ctx context.Context, product *model.ProductCategory) (string, error)
	GetProductCategories(ctx context.Context) ([]*model.ProductCategory, error)
	GetProductCategoryByID(ctx context.Context, productCategoryID string) (*model.ProductCategory, error)
}

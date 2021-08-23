package product

import (
	"context"
	"ecomjc-be/config"
	database "ecomjc-be/databases/interface"
	model "ecomjc-be/models"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *model.Product) (string, error)
	GetProducts(ctx context.Context, productCategories []*model.ProductCategory) ([]*model.Product, error)
	GetProductByID(ctx context.Context, productID string) (*model.Product, error)
	// GetProductByID(ctx context.Context, productID int) (*model.Product, error)
	// UpdateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	// DeleteProduct(ctx context.Context, productID int) (bool, error)

	CreateProductCategory(ctx context.Context, productCategory *model.ProductCategory) (string, error)
	GetProductCategories(ctx context.Context) ([]*model.ProductCategory, error)
	GetProductCategoryByID(ctx context.Context, productCategoryID string) (*model.ProductCategory, error)

	// CreateProductSpec(ctx context.Context, spec interface{}) (*model.Product, error)
	// GetProductSpecs(ctx context.Context) ([]interface{}, error)
	// GetProductSpecByID(ctx context.Context, specID string) (interface{}, error)
	// UpdateProductSpec(ctx context.Context, spec interface{}) (*model.Product, error)
	// DeleteProductSpec(ctx context.Context, specID string) (bool, error)

	// CreateProductPromotion(ctx context.Context, Product *model.Product) (*model.Product, error)
	// GetProductPromotion(ctx context.Context) ([]*model.Product, error)
	// GetProductPromotionByID(ctx context.Context, ProductID int) (*model.Product, error)
	// UpdateProductPromotion(ctx context.Context, Product *model.Product) (*model.Product, error)
	// DeleteProductPromotion(ctx context.Context, ProductID int) (bool, error)
}

type productService struct {
	conf     *config.Config
	database database.Interface
}

func NewProductService(conf *config.Config, database database.Interface) ProductService {
	return &productService{
		conf:     conf,
		database: database,
	}
}

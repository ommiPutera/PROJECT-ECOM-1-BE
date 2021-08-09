package endpoint

import (
	"context"
	service "ecomjc-be/services"
)

type Endpoint interface {
	// Products
	CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error)
	GetProducts(ctx context.Context, _ *GetProductsRequest) (*GetProductsResponse, error)
	GetProductByID(ctx context.Context, req *GetProductByIDRequest) (*GetProductByIDResponse, error)

	CreateProductCategory(ctx context.Context, req *CreateProductCategoryRequest) (*CreateProductCategoryResponse, error)
	GetProductCategories(ctx context.Context, _ *GetProductCategoriesRequest) (*GetProductCategoriesResponse, error)
	GetProductCategoryByID(ctx context.Context, req *GetProductCategoryByIDRequest) (*GetProductCategoryByIDResponse, error)
}

type endpoint struct {
	service *service.Service
}

func NewEndpoint(service *service.Service) Endpoint {
	return &endpoint{
		service: service,
	}
}

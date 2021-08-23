package endpoint

import (
	"context"
	model "ecomjc-be/models"
)

type CreateProductRequest struct {
	ID          string `json:"id"`
	ShopID      string `json:"shop_id"`
	CategoryID  string `json:"category_id"`
	SpecID      string `json:"spec_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
	Weight      uint   `json:"weight"`
}

type CreateProductResponse struct {
	ProductID string `json:"product_id"`
}

func (e *endpoint) CreateProduct(ctx context.Context, req *CreateProductRequest) (*CreateProductResponse, error) {
	product, err := e.service.Product.CreateProduct(ctx, &model.Product{
		ShopID:      req.ShopID,
		CategoryID:  req.CategoryID,
		SpecID:      req.SpecID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Weight:      req.Weight,
	})

	if err != nil {
		return nil, err
	}

	return &CreateProductResponse{
		ProductID: product,
	}, nil
}

type GetProductsRequest struct {
	AuthToken string `json:"-"`
}

type GetProductsResponse struct {
	Products []*model.Product `json:"products"`
}

func (e *endpoint) GetProducts(ctx context.Context, _ *GetProductsRequest) (*GetProductsResponse, error) {
	productCategories, err := e.service.Product.GetProductCategories(ctx)
	if err != nil {
		return nil, err
	}

	products, err := e.service.Product.GetProducts(ctx, productCategories)
	if err != nil {
		return nil, err
	}

	return &GetProductsResponse{
		Products: products,
	}, nil
}

type GetProductByIDRequest struct {
	AuthToken string `json:"-"`
	ProductID string
}

type GetProductByIDResponse struct {
	*model.Product `json:"product"`
}

func (e *endpoint) GetProductByID(ctx context.Context, req *GetProductByIDRequest) (*GetProductByIDResponse, error) {
	product, err := e.service.Product.GetProductByID(ctx, req.ProductID)
	if err != nil {
		return nil, err
	}

	productCategory, err := e.service.Product.GetProductCategoryByID(ctx, product.CategoryID)
	if err != nil {
		return nil, err
	}

	product.CategoryName = productCategory.Name

	return &GetProductByIDResponse{
		Product: product,
	}, nil
}

type CreateProductCategoryRequest struct {
	Name string `json:"category"`
}

type CreateProductCategoryResponse struct {
	ProductCategoryID string `json:"product_category_id"`
}

func (e *endpoint) CreateProductCategory(ctx context.Context, req *CreateProductCategoryRequest) (*CreateProductCategoryResponse, error) {
	productCategoryID, err := e.service.Product.CreateProductCategory(ctx, &model.ProductCategory{
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	return &CreateProductCategoryResponse{
		ProductCategoryID: productCategoryID,
	}, nil
}

type GetProductCategoriesRequest struct {
	AuthToken string `json:"-"`
}

type GetProductCategoriesResponse struct {
	ProductCategories []*model.ProductCategory `json:"product_categories"`
}

func (e *endpoint) GetProductCategories(ctx context.Context, _ *GetProductCategoriesRequest) (*GetProductCategoriesResponse, error) {
	productCategories, err := e.service.Product.GetProductCategories(ctx)
	if err != nil {
		return nil, err
	}

	return &GetProductCategoriesResponse{
		ProductCategories: productCategories,
	}, nil
}

type GetProductCategoryByIDRequest struct {
	AuthToken         string `json:"-"`
	ProductCategoryID string
}

type GetProductCategoryByIDResponse struct {
	ProductCategory *model.ProductCategory `json:"product_category"`
}

func (e *endpoint) GetProductCategoryByID(ctx context.Context, req *GetProductCategoryByIDRequest) (*GetProductCategoryByIDResponse, error) {
	productCategory, err := e.service.Product.GetProductCategoryByID(ctx, req.ProductCategoryID)
	if err != nil {
		return nil, err
	}

	return &GetProductCategoryByIDResponse{
		ProductCategory: productCategory,
	}, nil
}

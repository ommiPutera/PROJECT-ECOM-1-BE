package product

import (
	"context"
	model "ecomjc-be/models"
)

func (svc *productService) CreateProduct(ctx context.Context, product *model.Product) (string, error) {
	productID, err := svc.database.CreateProduct(ctx, product)
	if err != nil {
		return "", err
	}

	return productID, nil
}

func (svc *productService) GetProducts(ctx context.Context, productCategories []*model.ProductCategory) ([]*model.Product, error) {
	products, err := svc.database.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	categoryNames := make(map[string]string)

	for i := 0; i < len(productCategories); i += 1 {
		categoryNames[productCategories[i].ID] = productCategories[i].Name
	}

	for i := range products {
		products[i].CategoryName = categoryNames["ObjectID(\""+products[i].CategoryID+"\")"]
	}

	return products, nil
}

func (svc *productService) GetProductByID(ctx context.Context, productID string) (*model.Product, error) {
	product, err := svc.database.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (svc *productService) CreateProductCategory(ctx context.Context, productCategory *model.ProductCategory) (string, error) {
	productCategoryID, err := svc.database.CreateProductCategory(ctx, productCategory)
	if err != nil {
		return "", err
	}

	return productCategoryID, nil
}

func (svc *productService) GetProductCategories(ctx context.Context) ([]*model.ProductCategory, error) {
	productCategories, err := svc.database.GetProductCategories(ctx)
	if err != nil {
		return nil, err
	}

	return productCategories, nil
}

func (svc *productService) GetProductCategoryByID(ctx context.Context, productCategoryID string) (*model.ProductCategory, error) {
	productCategory, err := svc.database.GetProductCategoryByID(ctx, productCategoryID)
	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

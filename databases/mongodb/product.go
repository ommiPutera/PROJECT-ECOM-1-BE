package mongodbclient

import (
	"context"
	model "ecomjc-be/models"
	"encoding/json"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const productCollectionName string = "products"

func (db *mongodbClient) CreateProduct(ctx context.Context, product *model.Product) (string, error) {
	product.ID = primitive.NewObjectID().String()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result, err := db.client.Database("ecomjc").Collection(productCollectionName).InsertOne(ctx, product)
	if err != nil {
		return "", err
	}

	productID := strings.Replace(result.InsertedID.(string), "ObjectID(\"", "", -1)
	productID = strings.Replace(productID, "\")", "", -1)

	return productID, err
}

func (db *mongodbClient) GetProducts(ctx context.Context) ([]*model.Product, error) {
	cursor, err := db.client.Database("ecomjc").Collection(productCollectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var products []*model.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, err
}

func (db *mongodbClient) GetProductByID(ctx context.Context, productID string) (*model.Product, error) {
	productObjectID := "ObjectID(\"" + productID + "\")"

	var result bson.M
	err := db.client.Database("ecomjc").Collection(productCollectionName).FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: productObjectID}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	jsonbody, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	product := model.Product{}
	if err := json.Unmarshal(jsonbody, &product); err != nil {
		return nil, err
	}

	return &product, err
}

const productCategoryCollectionName string = "product_categories"

func (db *mongodbClient) CreateProductCategory(ctx context.Context, product *model.ProductCategory) (string, error) {
	product.ID = primitive.NewObjectID().String()

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", err
	}

	product.UpdatedAt = time.Now().In(location)

	result, err := db.client.Database("ecomjc").Collection(productCategoryCollectionName).InsertOne(ctx, product)
	if err != nil {
		return "", err
	}

	productCategoryID := strings.Replace(result.InsertedID.(string), "ObjectID(\"", "", -1)
	productCategoryID = strings.Replace(productCategoryID, "\")", "", -1)

	return productCategoryID, err
}

func (db *mongodbClient) GetProductCategories(ctx context.Context) ([]*model.ProductCategory, error) {
	cursor, err := db.client.Database("ecomjc").Collection(productCategoryCollectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var productCategories []*model.ProductCategory
	if err = cursor.All(ctx, &productCategories); err != nil {
		return nil, err
	}

	return productCategories, err
}

func (db *mongodbClient) GetProductCategoryByID(ctx context.Context, productCategoryID string) (*model.ProductCategory, error) {
	productCategoryObjectID := "ObjectID(\"" + productCategoryID + "\")"

	var result bson.M
	err := db.client.Database("ecomjc").Collection(productCategoryCollectionName).FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: productCategoryObjectID}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	jsonbody, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	productCategory := model.ProductCategory{}
	if err := json.Unmarshal(jsonbody, &productCategory); err != nil {
		return nil, err
	}

	productCategory.ID = productCategoryID

	return &productCategory, err
}

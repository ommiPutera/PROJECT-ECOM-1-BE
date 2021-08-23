package route

import (
	endpoint "ecomjc-be/endpoints"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (hr *HTTPRouter) InitPostCreateProduct() {
	hr.router.POST("/v1/products", func(resWriter http.ResponseWriter, httpReq *http.Request, _ httprouter.Params) {
		req := endpoint.CreateProductRequest{}

		decoder := json.NewDecoder(httpReq.Body)
		err := decoder.Decode(&req)
		if err != nil {
			EncodeErrorResponse(resWriter, err)
			return
		}

		res, err := hr.endpoint.CreateProduct(httpReq.Context(), &req)
		if err != nil {
			EncodeErrorResponse(resWriter, err)
			return
		}

		EncodeResponse(resWriter, res, http.StatusCreated)
	})
}

func (hr *HTTPRouter) InitGetProducts() {
	hr.router.GET("/v1/products", func(resWriter http.ResponseWriter, httpReq *http.Request, _ httprouter.Params) {
		res, err := hr.endpoint.GetProducts(httpReq.Context(), &endpoint.GetProductsRequest{})
		if err != nil {
			EncodeErrorResponse(resWriter, err)
			return
		}

		EncodeResponse(resWriter, res, http.StatusOK)
	})
}

func (hr *HTTPRouter) InitGetProductByID() {
	hr.router.GET("/v1/products/:product_id", func(resWriter http.ResponseWriter, httpReq *http.Request, param httprouter.Params) {
		req := endpoint.GetProductByIDRequest{}

		req.ProductID = param.ByName("product_id")

		res, err := hr.endpoint.GetProductByID(httpReq.Context(), &req)
		if err != nil {
			EncodeErrorResponse(resWriter, err)
			return
		}

		EncodeResponse(resWriter, res, http.StatusOK)
	})
}

func (hr *HTTPRouter) InitPostCreateProductCategory() {
	hr.router.POST("/v1/product-categories", func(resWriter http.ResponseWriter, httpReq *http.Request, _ httprouter.Params) {
		req := endpoint.CreateProductCategoryRequest{}

		decoder := json.NewDecoder(httpReq.Body)
		err := decoder.Decode(&req)
		if err != nil {
			EncodeErrorResponse(resWriter, err)
			return
		}

		res, err := hr.endpoint.CreateProductCategory(httpReq.Context(), &req)
		if err != nil {
			EncodeErrorResponse(resWriter, err)
			return
		}

		EncodeResponse(resWriter, res, http.StatusCreated)
	})
}

func (hr *HTTPRouter) InitGetProductCategories() {
	hr.router.GET("/v1/product-categories", func(resWriter http.ResponseWriter, httpReq *http.Request, _ httprouter.Params) {
		res, err := hr.endpoint.GetProductCategories(httpReq.Context(), &endpoint.GetProductCategoriesRequest{})
		if err != nil {
			EncodeErrorResponse(resWriter, err)
			return
		}

		EncodeResponse(resWriter, res, http.StatusOK)
	})
}

func (hr *HTTPRouter) InitGetProductCategoryByID() {
	hr.router.GET("/v1/product-categories/:product_category_id", func(resWriter http.ResponseWriter, httpReq *http.Request, param httprouter.Params) {
		req := endpoint.GetProductCategoryByIDRequest{}

		req.ProductCategoryID = param.ByName("product_category_id")

		res, err := hr.endpoint.GetProductCategoryByID(httpReq.Context(), &req)
		if err != nil {
			EncodeErrorResponse(resWriter, err)
			return
		}

		EncodeResponse(resWriter, res, http.StatusOK)
	})
}

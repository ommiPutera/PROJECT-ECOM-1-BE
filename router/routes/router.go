package route

import (
	endpoint "ecomjc-be/endpoints"

	"github.com/julienschmidt/httprouter"
)

type HTTPRouter struct {
	router   *httprouter.Router
	endpoint endpoint.Endpoint
}

func NewHTTPRouter(ep endpoint.Endpoint) *HTTPRouter {
	hr := &HTTPRouter{
		router:   httprouter.New(),
		endpoint: ep,
	}

	enableCORSSupport(hr.router)

	hr.InitPostCreateProduct()
	hr.InitGetProducts()
	hr.InitGetProductByID()

	hr.InitPostCreateProductCategory()
	hr.InitGetProductCategories()
	hr.InitGetProductCategoryByID()

	return hr
}

func (hr *HTTPRouter) GetRouter() *httprouter.Router {
	return hr.router
}

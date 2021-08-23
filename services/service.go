package service

import (
	"ecomjc-be/config"
	mongodbclient "ecomjc-be/databases/mongodb"
	"ecomjc-be/services/address"
	"ecomjc-be/services/cart"
	"ecomjc-be/services/chat"
	"ecomjc-be/services/ewallet"
	"ecomjc-be/services/feedback"
	"ecomjc-be/services/order"
	"ecomjc-be/services/product"
	"ecomjc-be/services/shop"
	"ecomjc-be/services/transaction"
	"ecomjc-be/services/user"
)

type Service struct {
	Auth        user.UserService
	Address     address.AddressService
	Cart        cart.CartService
	Chat        chat.ChatService
	Ewallet     ewallet.EWalletService
	Feedback    feedback.FeedbackService
	Order       order.OrderService
	Product     product.ProductService
	Shop        shop.ShopService
	Transaction transaction.TransactionService
}

func NewService(conf *config.Config) *Service {
	mongodbClient := mongodbclient.NewMongoDBClient(conf)

	productService := product.NewProductService(conf, mongodbClient)

	return &Service{
		Product: productService,
	}
}

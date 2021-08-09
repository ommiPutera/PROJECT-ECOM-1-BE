package server

import (
	"ecomjc-be/config"
	endpoint "ecomjc-be/endpoints"
	route "ecomjc-be/router/routes"
	service "ecomjc-be/services"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func ServeHTTP(conf *config.Config) *http.Server {
	svc := service.NewService(conf)

	ep := endpoint.NewEndpoint(svc)

	r := route.NewHTTPRouter(ep)

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"Authorization"},
	})

	return &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.HTTPPort),
		Handler: c.Handler(r.GetRouter()),
	}
}

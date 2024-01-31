package main

import (
	"log"

	"product-catalog-manager/internal/configuration"
	"product-catalog-manager/internal/dependency_provider"
	"product-catalog-manager/internal/product_catalog"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	dp := dependency_provider.New(config)
	product_catalog.RegisterRouter(app, dp.GetProductService())
	log.Fatal(app.Listen(":" + config.Port))
}

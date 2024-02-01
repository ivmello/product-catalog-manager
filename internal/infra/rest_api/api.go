package rest_api

import (
	"log"

	"product-catalog-manager/internal/dependency_provider"
	"product-catalog-manager/internal/product_catalog"

	"github.com/gofiber/fiber/v2"
)

func InitializeServer(dp *dependency_provider.DependencyProvider) {
	app := fiber.New()
	product_catalog.RegisterRouter(app, dp.GetProductService())
	log.Fatal(app.Listen(":" + dp.GetConfig().Port))
}

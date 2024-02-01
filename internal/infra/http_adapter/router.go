package http_adapter

import (
	"product-catalog-manager/internal/application/dependency_provider"
	"product-catalog-manager/internal/application/product"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(app *fiber.App, db *dependency_provider.DependencyProvider) {
	productHandler := product.NewProductHandler(db.GetProductService())
	productGroup := app.Group("products")
	productGroup.Post("/", productHandler.CreateProduct)
	productGroup.Get("/", productHandler.ListProducts)
}

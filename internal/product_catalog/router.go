package product_catalog

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(app *fiber.App, service ProductService) {
	handler := NewProductHandler(service)
	app.Post("/products", handler.CreateProduct)
	app.Get("/products", handler.ListProducts)
}

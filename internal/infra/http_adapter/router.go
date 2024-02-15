package http_adapter

import (
	"product-catalog-manager/internal/application/dependency_provider"
	"product-catalog-manager/internal/application/product"
	"product-catalog-manager/internal/infra/websocket_adapter"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(app *fiber.App, dp *dependency_provider.DependencyProvider, msgChan chan []byte) {
	productHandler := product.NewProductHandler(dp.GetProductService())
	websocketHandler := websocket_adapter.New(app)
	productGroup := app.Group("products")
	productGroup.Post("/", productHandler.CreateProduct)
	productGroup.Get("/", productHandler.ListProducts)
	app.Get("/ws/products", websocket.New(func(c *websocket.Conn) {
		websocketHandler.Listener(c, msgChan)
	}))
}

package http_adapter

import (
	"log"

	"product-catalog-manager/internal/application/dependency_provider"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func InitializeServer(dp *dependency_provider.DependencyProvider, msgChan chan []byte) {
	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	RegisterRouter(app, dp, msgChan)
	log.Fatal(app.Listen(":" + dp.GetConfig().Port))
}

package http

import (
	"log"

	"product-catalog-manager/internal/application/dependency_provider"

	"github.com/gofiber/fiber/v2"
)

func InitializeServer(dp *dependency_provider.DependencyProvider) {
	app := fiber.New()
	RegisterRouter(app, dp)
	log.Fatal(app.Listen(":" + dp.GetConfig().Port))
}

package http_adapter

import (
	"log"

	"product-catalog-manager/internal/application/dependency_provider"
	"product-catalog-manager/internal/application/product"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type client struct{}

var (
	clients    = make(map[*websocket.Conn]client)
	register   = make(chan *websocket.Conn)
	broadcast  = make(chan string)
	unregister = make(chan *websocket.Conn)
)

func RegisterRouter(app *fiber.App, dp *dependency_provider.DependencyProvider, msgChan <-chan []byte) {
	productHandler := product.NewProductHandler(dp.GetProductService())
	productGroup := app.Group("products")
	productGroup.Post("/", productHandler.CreateProduct)
	productGroup.Get("/", productHandler.ListProducts)
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	go websocketHandler(dp, msgChan)
	app.Get("/ws/products", websocket.New(func(c *websocket.Conn) {
		defer func() {
			unregister <- c
			c.Close()
		}()
		register <- c
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}
				return
			}
			broadcast <- string(message)
		}
	}))
}

func websocketHandler(dp *dependency_provider.DependencyProvider, msgChan <-chan []byte) {
	for {
		select {
		case connection := <-register:
			clients[connection] = client{}
			log.Println("connection registered:", clients[connection])
		case message := <-broadcast:
			log.Println("message received:", string(message))
			for connection := range clients {
				dp.GetProductService().SendMessage([]byte(message))
				if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Println("write error:", err)
					unregister <- connection
					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
				}
			}
		case message := <-msgChan:
			log.Println("message received:", string(message))
			for connection := range clients {
				if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Println("write error:", err)
					unregister <- connection
					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
				}
			}
		case connection := <-unregister:
			delete(clients, connection)
			log.Println("connection unregistered")
		}
	}
}

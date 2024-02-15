package websocket_adapter

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type WebsocketHandler interface {
	Listener(c *websocket.Conn, msgChan chan []byte) error
}

type handler struct {
	app *fiber.App
}

func New(app *fiber.App) WebsocketHandler {
	return &handler{
		app: app,
	}
}

func (h *handler) Listener(c *websocket.Conn, msgChan chan []byte) error {
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		msgChan <- msg
		if err = c.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
	return nil
}

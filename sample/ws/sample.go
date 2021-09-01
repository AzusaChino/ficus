package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

func _() {
	app := fiber.New()
	// 1. enable ws middleware
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// 2. sample ws usage
	app.Get("/ws/:id", websocket.New(func(conn *websocket.Conn) {
		log.Println(conn.Locals("allowed"))
		log.Println(conn.Params("id"))
		log.Println(conn.Query("v"))
		log.Println(conn.Cookies("session"))

		var (
			mt  int
			msg []byte
			err error
		)
		// for loop for listening
		for {
			if mt, msg, err = conn.ReadMessage(); err != nil {
				log.Fatalf("read error: %v\n", err)
			}
			log.Printf("recv msg: %s", msg)
			if err = conn.WriteMessage(mt, msg); err != nil {
				log.Fatalf("write error: %v\n", err)
			}
		}
	}))
}

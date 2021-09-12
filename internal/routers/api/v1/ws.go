package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

func WsHandler() fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
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
	})
}

package v1

import (
	"github.com/azusachino/ficus/global"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func WsHandler() fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)
		// for loop for listening
		for {
			if mt, msg, err = conn.ReadMessage(); err != nil {
				global.Logger.Fatalf("read error: %v\n", err)
			}
			global.Logger.Infof("recv msg: %s", msg)
			if err = conn.WriteMessage(mt, msg); err != nil {
				global.Logger.Fatalf("write error: %v\n", err)
			}
		}
	})
}

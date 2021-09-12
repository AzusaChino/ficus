package routers

import (
	"github.com/AzusaChino/ficus/internal/middleware/ws"
	"github.com/AzusaChino/ficus/internal/routers/api/v1"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {

	// normal http router
	apiV1 := app.Group("/api/v1")
	{
		apiV1.Post("/logging/collect", v1.UploadFile)
		apiV1.Post("/sample/hello", v1.Hello)
	}

	// ws router
	app.Use("/ws", ws.New())
	app.Get("/ws/:id", v1.WsHandler())
}

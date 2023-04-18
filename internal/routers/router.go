package routers

import (
	"github.com/azusachino/ficus/internal/middleware/ws"
	v1 "github.com/azusachino/ficus/internal/routers/api/v1"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {

	// normal http router
	apiV1 := app.Group("/api/v1")
	{
		apiV1.Post("/logging/collect", v1.UploadFile)
		apiV1.Get("/hello/:person", v1.Hello)
		apiV1.Post("/hi/:person", v1.Hello)
		apiV1.Get("/say/:name", v1.SayWhat)

		apiV1Grpc := apiV1.Group("/grpc")
		{
			apiV1Grpc.Get("/hello/:msg", v1.SayHello)
		}
	}

	// ws router
	app.Use("/ws", ws.New())
	app.Get("/ws/:id", v1.WsHandler())
}

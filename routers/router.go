package routers

import (
	v1 "github.com/AzusaChino/ficus/routers/api/v1"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {

	apiV1 := app.Group("/api/v1")
	{
		apiV1.Post("/logging/collect", v1.UploadFile)
		apiV1.Post("/sample/hello", v1.Hello)
	}

}

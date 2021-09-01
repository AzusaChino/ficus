package routers

import (
	v12 "github.com/AzusaChino/ficus/internal/routers/api/v1"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {

	apiV1 := app.Group("/api/v1")
	{
		apiV1.Post("/logging/collect", v12.UploadFile)
		apiV1.Post("/sample/hello", v12.Hello)
	}

}

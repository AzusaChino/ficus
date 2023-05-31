package routers

import (
	v1 "github.com/azusachino/ficus/internal/routers/api/v1"
	fa "github.com/azusachino/ficus/pkg/app"
	"github.com/azusachino/ficus/pkg/errcode"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {

	apiV1 := app.Group("/api/v1")

	{
		apiV1.Get("/ping", func(ctx *fiber.Ctx) error {
			fa.NewResponse(ctx).ToErrorResponse(errcode.NotFound)
			return nil
		})
        
		// Tag related routers
		tag := v1.NewTag()
		apiV1.Get("/tags", tag.List)
		apiV1.Post("/tags", tag.Create)
		apiV1.Put("/tags/:id", tag.Update)
		apiV1.Patch("/tags/:id/state", tag.Update)
		apiV1.Delete("/tags/:id", tag.Delete)

		// Article related routers
		article := v1.NewArticle()
		apiV1.Get("/articles/:id", article.Get)
		apiV1.Get("/articles", article.List)
		apiV1.Post("/articles", article.Create)
		apiV1.Put("/articles/:id", article.Update)
		apiV1.Patch("/articles/:id/state", article.Update)
		apiV1.Delete("/articles/:id", article.Delete)
	}

}

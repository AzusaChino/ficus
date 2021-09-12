package view

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"io"
)

type errorTemplateEngine struct{}

func (t errorTemplateEngine) Render(w io.Writer, name string, bind interface{}, layout ...string) error {
	return errors.New("errorTemplateEngine")
}

func (t errorTemplateEngine) Load() error { return nil }

func _(app *fiber.App) {

	app = fiber.New(fiber.Config{
		Views: errorTemplateEngine{},
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("home", fiber.Map{
			"title": "HomePage",
			"year":  2021,
		})
	})
}

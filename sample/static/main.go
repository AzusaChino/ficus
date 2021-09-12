package static

import "github.com/gofiber/fiber/v2"

func _(app *fiber.App) {
	app.Static("/", "./public")
	app.Static("/prefix", "./public")
	app.Static("/*", "./public/index.html")

}

package main

import (
	"fmt"
	"github.com/AzusaChino/ficus/pkg/conf"
	"github.com/AzusaChino/ficus/pkg/kafka"
	"github.com/AzusaChino/ficus/pkg/logging"
	"github.com/AzusaChino/ficus/pkg/pool"
	"github.com/AzusaChino/ficus/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"net/http"
)

func init() {
	conf.Setup()
	logging.Setup()
	kafka.Setup()
	pool.Setup()
}

func main() {
	defer pool.Pool.Release()
	defer kafka.Close()

	cnf := fiber.Config{
		ReadTimeout:  conf.ServerConfig.ReadTimeout,
		WriteTimeout: conf.ServerConfig.WriteTimeout,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusInternalServerError).JSON(fmt.Sprintf(`{"error":%v}`, err))
		},
	}
	app := fiber.New(cnf)
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(recover.New())

	// first append url, second local folder
	app.Static("/static", "./static")
	routers.InitRouter(app)
	endPoint := fmt.Sprintf(":%d", conf.ServerConfig.HttpPort)

	log.Printf("[info] start http server listening %s", endPoint)
	_ = app.Listen(endPoint)
}

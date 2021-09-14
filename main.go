package main

import (
	"fmt"
	"github.com/AzusaChino/ficus/internal/middleware/fiberprometheus"
	"github.com/AzusaChino/ficus/internal/middleware/fibertracing"
	"github.com/AzusaChino/ficus/internal/routers"
	"github.com/AzusaChino/ficus/pkg/conf"
	"github.com/AzusaChino/ficus/pkg/kafka"
	"github.com/AzusaChino/ficus/pkg/logging"
	"github.com/AzusaChino/ficus/pkg/pool"
	"github.com/AzusaChino/ficus/pkg/tracer"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	app.Use(logger.New())
	app.Use(recover.New())
	tracer.New(tracer.Config{
		ServiceName: appName,
	})
	app.Use(fibertracing.New(fibertracing.Config{
		Tracer: opentracing.GlobalTracer(),
	}))

	// prometheus metric
	prometheus := fiberprometheus.New(appName)
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Do)

	// first append url, second local folder
	app.Static("/static", "./static")
	routers.InitRouter(app)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})
	endPoint := fmt.Sprintf(":%d", conf.ServerConfig.HttpPort)
	log.Printf("[info] start http server listening %s", endPoint)

	go func() {
		if err := app.Listen(endPoint); err != nil && app.Server() != nil {
			log.Fatalf("app.Listen error: %v\n", err)
		}
	}()

	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign

	log.Println("shutting down the server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("app shutdown error: %v\n", err)
	}
	log.Println("server shut down finished.")
}

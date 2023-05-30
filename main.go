package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/azusachino/ficus/internal/middleware/fiberprometheus"
	"github.com/azusachino/ficus/internal/middleware/fibertracing"
	"github.com/azusachino/ficus/internal/routers"
	"github.com/azusachino/ficus/pkg/conf"
	"github.com/azusachino/ficus/pkg/logging"
	"github.com/azusachino/ficus/pkg/mydb"
	"github.com/azusachino/ficus/pkg/pool"
	"github.com/azusachino/ficus/pkg/tracer"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/opentracing/opentracing-go"
)

const appName = "ficus"

// Init all necessary components
func init() {
	logging.Setup()
	pool.Setup()
	mydb.SetUp()
}

func main() {
	defer pool.Close()
	defer mydb.Close()

	serverConfig := conf.Config.Server
	cnf := fiber.Config{
		ReadTimeout:  serverConfig.ReadTimeout,
		WriteTimeout: serverConfig.WriteTimeout,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusInternalServerError).JSON(fmt.Sprintf(`{"error":%v}`, err))
		},
		AppName: appName,
	}

	app := fiber.New(cnf)
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(logger.New(
		logger.Config{
			Format:     "[${time}] ${pid} ${status} - ${method} ${path} ${latency}\n",
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Asia/Shanghai",
		}))
	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		}))
	tracer.New(tracer.Config{
		ServiceName: appName,
	})
	app.Use(fibertracing.New(fibertracing.Config{
		Tracer: opentracing.GlobalTracer(),
	}))

	routers.InitRouter(app)

	// prometheus metric
	prometheus := fiberprometheus.New(appName)
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Handler())

	// first append url, second local folder
	app.Static("/static", "./static")

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})

	endPoint := fmt.Sprintf(":%d", serverConfig.HttpPort)
	log.Printf("start http server listening %s\n", endPoint)

	go func() {
		if err := app.Listen(endPoint); err != nil && app.Server() != nil {
			log.Fatalf("app.Listen error: %v\n", err)
		}
	}()

	// listen on INT/TERM
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign

	log.Println("shutting down the server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("app shutdown error: %v\n", err)
	}

	log.Println("server shut down finished")
}

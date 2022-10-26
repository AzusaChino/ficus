package main

import (
	"fmt"
	"github.com/azusachino/ficus/internal/middleware/fiberprometheus"
	"github.com/azusachino/ficus/internal/middleware/fibertracing"
	"github.com/azusachino/ficus/internal/routers"
	"github.com/azusachino/ficus/pkg/conf"
	"github.com/azusachino/ficus/pkg/etcd"
	"github.com/azusachino/ficus/pkg/kafka"
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
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const appName = "ficus"

func init() {
	conf.Setup()
	logging.Setup()
	pool.Setup()
	kafka.Setup()
	mydb.SetUp()
	etcd.Setup()
}

func main() {
	defer pool.Close()
	defer kafka.Close()
	defer mydb.Close()
	defer etcd.Close()
	cnf := fiber.Config{
		ReadTimeout:  conf.ServerConfig.ReadTimeout,
		WriteTimeout: conf.ServerConfig.WriteTimeout,
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
	app.Use(prometheus.Do())

	// first append url, second local folder
	app.Static("/static", "./static")

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})
	endPoint := fmt.Sprintf(":%d", conf.ServerConfig.HttpPort)
	log.Printf("start http server listening %s", endPoint)

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

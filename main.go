package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/azusachino/ficus/global"
	"github.com/azusachino/ficus/internal/dao"
	"github.com/azusachino/ficus/internal/middleware/fiberprometheus"
	"github.com/azusachino/ficus/internal/routers"
	fl "github.com/azusachino/ficus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

const appName = "ficus"

func init() {
	// load env
	godotenv.Load("ficus.env")

	// 1. set up global config
	initConfig()
	// 2. setup logger
	initLogger()
	// 3. working pool
	initWorkingPool()
	// 4. init database
	initDb()
}

func main() {
	// ensure resources are released
	defer func() {
		if global.Pool != nil {
			global.Pool.Release()
		}
	}()

	// start the ficus server
	serverConfig := global.Config.Server
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
	// tracer.New(tracer.Config{
	// 	ServiceName: appName,
	// })
	// app.Use(fibertracing.New(fibertracing.Config{
	// 	Tracer: opentracing.GlobalTracer(),
	// }))

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
	global.Logger.Infof("start http server listening %s\n", endPoint)

	go func() {
		if err := app.Listen(endPoint); err != nil && app.Server() != nil {
			global.Logger.Fatalf("app.Listen error: %v\n", err)
		}
	}()

	// listen on INT/TERM
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign

	global.Logger.Info("shutting down the server...")

	if err := app.Shutdown(); err != nil {
		global.Logger.Fatalf("app shutdown error: %v\n", err)
	}

	global.Logger.Info("server shut down finished")
}

func initConfig() {
	var err error
	vp := viper.New()
	vp.SetConfigName("ficus")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("configs")
	if err = vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error when reading config file: %w", err))
	}
	err = vp.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("fatal error when unmarshal config file: %w", err))
	}
}

func initWorkingPool() {
	var err error
	size := runtime.NumCPU()

	global.Pool, err = ants.NewPool(size,
		ants.WithExpiryDuration(100*time.Second),
		ants.WithPanicHandler(func(i interface{}) {
			global.Logger.Fatal(i)
		}),
		ants.WithLogger(log.Default()))
	if err != nil {
		panic(fmt.Errorf("error when setup ants pool: %v", err))
	}
}

func initDb() {
	var err error
	global.DbEngine, err = dao.NewDbEngine(&global.Config.Database)
	if err != nil {
		panic(fmt.Errorf("fatal error when connect database: %w", err))
	}
}

func initLogger() {
	appConfig := global.Config.App
	logFileName := fmt.Sprintf("%s%s%s-%s.%s",
		appConfig.LogFileLocation,
		string(os.PathSeparator),
		appConfig.LogFileSaveName,
		time.Now().Format(appConfig.TimeFormat),
		appConfig.LogFileExt,
	)
	lj := lumberjack.Logger{
		Filename:  logFileName,
		MaxSize:   600,
		MaxAge:    7,
		LocalTime: true,
	}
	global.Logger = fl.NewLogger(&lj, "", log.LstdFlags).WithCaller(2)
}

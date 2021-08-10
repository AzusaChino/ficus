package main

import (
	"fmt"
	"github.com/AzusaChino/ficus/pkg/conf"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.SugaredLogger

func init() {
	conf.Setup()
	lg, _ := zap.NewProduction()
	defer lg.Sync()
	logger = lg.Sugar()
}

func main() {
	gin.SetMode(conf.ServerConfig.RunMode)
	fmt.Println(conf.KafkaConfig.Host)
	fmt.Println(conf.AppConfig.LogFileLocation)

	server := &http.Server{}

	logger.Infof("server start http server listening %s", "8080")
	_ = server.ListenAndServe()
}

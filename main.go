package main

import (
	"fmt"
	"github.com/AzusaChino/ficus/pkg/conf"
	"github.com/AzusaChino/ficus/pkg/kafka"
	"github.com/AzusaChino/ficus/pkg/logging"
	"github.com/AzusaChino/ficus/pkg/pool"
	"github.com/AzusaChino/ficus/routers"
	"github.com/gin-gonic/gin"
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
	gin.SetMode(conf.ServerConfig.RunMode)

	router := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", conf.ServerConfig.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    conf.ServerConfig.ReadTimeout,
		WriteTimeout:   conf.ServerConfig.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()

}

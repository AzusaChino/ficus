package main

import (
	"github.com/AzusaChino/ficus/pkg/conf"
	"github.com/gin-gonic/gin"
)

func init() {
	conf.Setup()
}

func main() {
	gin.SetMode(conf.ServerConfig.RunMode)


}

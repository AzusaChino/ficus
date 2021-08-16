package routers

import (
	v1 "github.com/AzusaChino/ficus/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/logging/collect", v1.UploadFile)
		apiV1.POST("/sample/hello", v1.Hello)
	}
	return r
}

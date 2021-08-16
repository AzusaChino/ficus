package v1

import (
	"github.com/AzusaChino/ficus/pkg/pool"
	"github.com/gin-gonic/gin"
	"log"
)

type HelloParam struct {
	Person string `json:"person"`
}

func Hello(c *gin.Context) {
	var param = HelloParam{}
	_ = c.ShouldBindJSON(&param)
	log.Printf("[info] hello start")
	_ = pool.Pool.Submit(func() {
		log.Printf("[info] hello, %s\n", param.Person)
	})
	log.Printf("[info] hello end")
	c.JSON(200, "OK")
}

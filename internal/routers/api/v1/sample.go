package v1

import (
	"github.com/AzusaChino/ficus/pkg/pool"
	"github.com/gofiber/fiber/v2"
	"log"
)

type HelloParam struct {
	Person string `json:"person"`
}

func Hello(c *fiber.Ctx) error {
	var param = HelloParam{}
	_ = c.BodyParser(&param)
	log.Printf("[info] hello start")
	_ = pool.Pool.Submit(func() {
		log.Printf("[info] hello, %s\n", param.Person)
	})
	log.Printf("[info] hello end")
	return c.SendString("ok")
}

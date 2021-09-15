package v1

import (
	"fmt"
	"github.com/AzusaChino/ficus/pkg/pool"
	"github.com/AzusaChino/ficus/service/grpc_service"
	"github.com/gofiber/fiber/v2"
	"log"
)

type HelloParam struct {
	Person string `json:"person"`
}

func Hello(c *fiber.Ctx) error {
	person := c.Params("person")
	msg := fmt.Sprintf("hello, %s", person)
	_ = pool.Pool.Submit(func() {
		log.Println(msg)
	})
	return c.SendString(msg)
}

func SayHello(ctx *fiber.Ctx) error {
	msg := ctx.Params("msg")
	fastCtx := ctx.Context()
	r := grpc_service.DoHello(msg, fastCtx)
	return ctx.SendString(r)
}

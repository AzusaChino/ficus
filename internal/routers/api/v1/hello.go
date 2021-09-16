package v1

import (
	"context"
	"fmt"
	"github.com/AzusaChino/ficus/pkg/pool"
	"github.com/AzusaChino/ficus/service/grpc_service"
	"github.com/AzusaChino/ficus/service/my_service"
	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
)

var (
	mysqlTag = opentracing.Tag{
		Key:   string(ext.Component),
		Value: "mysql",
	}
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

func SayWhat(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	parentSpanCtx := opentracing.SpanFromContext(ctx.Context().Value("ctx").(context.Context)).Context()
	opts := []opentracing.StartSpanOption{
		opentracing.ChildOf(parentSpanCtx),
		ext.SpanKindConsumer,
		mysqlTag,
	}
	span := opentracing.GlobalTracer().StartSpan("SayWhat", opts...)
	msg := my_service.GetMsg(name)
	span.Finish()

	return ctx.SendString(msg)
}

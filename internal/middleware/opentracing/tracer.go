package opentracing

import (
	"fmt"
	"github.com/AzusaChino/ficus/global"
	"github.com/AzusaChino/ficus/util"
	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

const DefaultParentSpanKey = "#defaultParentSpanKey"

func Tracing() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var span opentracing.Span
		operationName := generateOperationName(c)
		hdr := make(http.Header)
		// no container for header, but we can visit them like below
		c.Request().Header.VisitAll(func(key, value []byte) {
			hdr.Set(util.GetString(key), util.GetString(value))
		})

		if spanCtx, err := global.Tracer.Extract(opentracing.HTTPHeaders, hdr); err != nil {
			span = global.Tracer.StartSpan(operationName)
		} else {
			span = global.Tracer.StartSpan(operationName, ext.RPCServerOption(spanCtx))
		}

		modify(c, span)

		defer func() {
			status := c.Response().StatusCode()
			ext.HTTPStatusCode.Set(span, uint16(status))
			if status >= fiber.StatusInternalServerError {
				ext.Error.Set(span, true)
			}
			span.Finish()
		}()

		c.Locals(DefaultParentSpanKey, span)

		return c.Next()
	}
}

func generateOperationName(c *fiber.Ctx) string {
	return fmt.Sprintf(`HTTP %s URL: %s`, c.Method(), c.Path())
}

func modify(ctx *fiber.Ctx, span opentracing.Span) {
	span.SetTag("http.remote_addr", ctx.IP())
	span.SetTag("http.path", ctx.Path())
	span.SetTag("http.host", ctx.Hostname())
	span.SetTag("http.method", ctx.Method())
	span.SetTag("http.url", ctx.OriginalURL())
}

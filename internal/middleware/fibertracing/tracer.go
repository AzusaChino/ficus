package fibertracing

import (
	"github.com/AzusaChino/ficus/util"
	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

func New(config Config) func(c *fiber.Ctx) error {
	cfg := configDefault(config)
	return func(c *fiber.Ctx) error {
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		var span opentracing.Span
		operationName := cfg.TransactionName(c)
		tracer := cfg.Tracer
		hdr := make(http.Header)
		// no container for header, but we can visit them like below
		c.Request().Header.VisitAll(func(key, value []byte) {
			hdr.Set(util.GetString(key), util.GetString(value))
		})

		if spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(hdr)); err != nil {
			span = tracer.StartSpan(operationName)
		} else {
			span = tracer.StartSpan(operationName, ext.RPCServerOption(spanCtx))
		}

		cfg.Modify(c, span)

		defer func() {
			status := c.Response().StatusCode()
			ext.HTTPStatusCode.Set(span, uint16(status))
			if status >= fiber.StatusInternalServerError {
				ext.Error.Set(span, true)
			}
			span.Finish()
		}()
		return c.Next()
	}
}

package fibertracing

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
)

type Config struct {
	Tracer          opentracing.Tracer
	TransactionName func(ctx *fiber.Ctx) string
	Filter          func(ctx *fiber.Ctx) bool
	Modify          func(ctx *fiber.Ctx, span opentracing.Span)
}

var DefaultConfig = Config{
	Tracer: opentracing.NoopTracer{},
	Modify: func(ctx *fiber.Ctx, span opentracing.Span) {
		span.SetTag("http.method", ctx.Method()) // GET, POST
		span.SetTag("http.remote_address", ctx.IP())
		span.SetTag("http.path", ctx.Path())
		span.SetTag("http.host", ctx.Hostname())
		span.SetTag("http.url", ctx.OriginalURL())
	},
	TransactionName: func(ctx *fiber.Ctx) string {
		return fmt.Sprintf(`HTTP %s URL: %s`, ctx.Method(), ctx.Path())
	},
}

// configDefault function to return default values
func configDefault(config ...Config) Config {
	// Return default config if no config provided
	if len(config) < 1 {
		return DefaultConfig
	}
	cfg := config[0]

	if cfg.Tracer == nil {
		cfg.Tracer = DefaultConfig.Tracer
	}

	if cfg.TransactionName == nil {
		cfg.TransactionName = DefaultConfig.TransactionName
	}

	if cfg.Modify == nil {
		cfg.Modify = DefaultConfig.Modify
	}

	return cfg
}

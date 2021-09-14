package tracer

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func New(c Config) {
	cfg := configDefault(c)
	tracer := InitJaeger(cfg)
	opentracing.SetGlobalTracer(tracer)
	return
}

func InitJaeger(c Config) opentracing.Tracer {
	cfg := config.Configuration{
		ServiceName: c.ServiceName,
		Sampler:     c.Sampler,
		Reporter:    c.Reporter,
		RPCMetrics:  c.EnableRpcMetrics,
		Headers:     c.Headers,
		Tags:        c.tags,
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		if c.PanicOnError {
			panic("Init jaeger failed")
		} else {
			fmt.Println("init jaeger failed")
		}
	}
	c.closer = closer.Close
	return tracer
}

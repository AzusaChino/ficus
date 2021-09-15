package tracer

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"os"
	"time"
)

type Config struct {
	ServiceName      string
	Sampler          *config.SamplerConfig
	Reporter         *config.ReporterConfig
	Headers          *jaeger.HeadersConfig
	EnableRpcMetrics bool
	tags             []opentracing.Tag
	options          []config.Option
	PanicOnError     bool
	closer           func() error
}

var DefaultConfig = Config{
	ServiceName: "default",
	Sampler: &config.SamplerConfig{
		Type:  "const",
		Param: 1,
	},
	Reporter: &config.ReporterConfig{
		LogSpans:            false,
		BufferFlushInterval: 1 * time.Second,
		LocalAgentHostPort:  fmt.Sprintf("%s:%d", jaeger.DefaultUDPSpanServerHost, jaeger.DefaultUDPSpanServerPort),
	},
	EnableRpcMetrics: true,
	Headers: &jaeger.HeadersConfig{
		TraceBaggageHeaderPrefix: jaeger.TraceBaggageHeaderPrefix,
		TraceContextHeaderName:   jaeger.TraceContextHeaderName,
	},
	tags: []opentracing.Tag{
		{Key: "hostname", Value: "hostname"},
	},
	PanicOnError: true,
}

// configDefault return default config of jaeger
func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return DefaultConfig
	}
	cfg := config[0]
	if addr := os.Getenv("JAEGER_AGENT_ADDR"); addr != "" {
		cfg.Reporter.LocalAgentHostPort = addr
	}

	if cfg.ServiceName == "" {
		cfg.ServiceName = DefaultConfig.ServiceName
	}

	if cfg.Sampler == nil {
		cfg.Sampler = DefaultConfig.Sampler
	}

	if cfg.Reporter == nil {
		cfg.Reporter = DefaultConfig.Reporter
	}

	if cfg.Headers == nil {
		cfg.Headers = DefaultConfig.Headers
	}

	if cfg.tags == nil {
		cfg.tags = DefaultConfig.tags
	}

	return cfg
}

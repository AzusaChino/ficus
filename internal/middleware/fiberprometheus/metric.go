package fiberprometheus

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
	"time"
)

type FiberPrometheus struct {
	requestTotal    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	requestInflight *prometheus.GaugeVec
	defaultUrl      string
}

func create(serviceName, namespace, subSystem string, labels map[string]string) *FiberPrometheus {
	constLabels := make(prometheus.Labels)
	if serviceName != "" {
		constLabels["service"] = serviceName
	}
	for l, v := range labels {
		constLabels[l] = v
	}

	counter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name:        prometheus.BuildFQName(namespace, subSystem, "request_total"),
			Help:        "Count all http requests by status code, method and path",
			ConstLabels: constLabels,
		},
		[]string{"status_code", "method", "path"})

	histogram := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        prometheus.BuildFQName(namespace, subSystem, "request_duration_seconds"),
			Help:        "Duration of all http requests by status_code, method and path",
			ConstLabels: constLabels,
			Buckets: []float64{
				0.000000001, // 1ns
				0.000000002,
				0.000000005,
				0.00000001, // 10ns
				0.00000002,
				0.00000005,
				0.0000001, // 100ns
				0.0000002,
				0.0000005,
				0.000001, // 1µs
				0.000002,
				0.000005,
				0.00001, // 10µs
				0.00002,
				0.00005,
				0.0001, // 100µs
				0.0002,
				0.0005,
				0.001, // 1ms
				0.002,
				0.005,
				0.01, // 10ms
				0.02,
				0.05,
				0.1, // 100 ms
				0.2,
				0.5,
				1.0, // 1s
				2.0,
				5.0,
				10.0, // 10s
				15.0,
				20.0,
				30.0,
			},
		},
		[]string{"status_code", "method", "path"})

	gauge := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name:        prometheus.BuildFQName(namespace, subSystem, "requests_in_progress_total"),
		Help:        "All the requests in progress",
		ConstLabels: constLabels,
	}, []string{"method", "path"})

	return &FiberPrometheus{
		requestTotal:    counter,
		requestDuration: histogram,
		requestInflight: gauge,
		defaultUrl:      "/metrics",
	}
}

func New(serviceName string) *FiberPrometheus {
	return create(serviceName, "http", "", nil)
}

func NewWith(serviceName, namespace, subsystem string) *FiberPrometheus {
	return create(serviceName, namespace, subsystem, nil)
}

func NewWithLabels(labels map[string]string, namespace, subsystem string) *FiberPrometheus {
	return create("", namespace, subsystem, labels)
}

func (ps *FiberPrometheus) RegisterAt(app *fiber.App, url string) {
	ps.defaultUrl = url
	app.Get(ps.defaultUrl, adaptor.HTTPHandler(promhttp.Handler()))
}

func (ps *FiberPrometheus) Do() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()
		r := ctx.Route()
		method := r.Method
		path := r.Path

		if path == ps.defaultUrl {
			return ctx.Next()
		}

		ps.requestInflight.WithLabelValues(method, path).Inc()
		defer func() {
			ps.requestInflight.WithLabelValues(method, path).Dec()
		}()

		statusCode := strconv.Itoa(ctx.Response().StatusCode())
		ps.requestTotal.WithLabelValues(statusCode, method, path).Inc()

		elapsed := float64(time.Since(start).Nanoseconds()) / 1000000000
		ps.requestDuration.WithLabelValues(statusCode, method, path).Observe(elapsed)
		return nil
	}
}

// Package telemetry sets up OpenTelemetry traces, metrics and logs for every service.
//
// Metrics are always served in Prometheus format on METRICS_ADDR (default :9464,
// path /metrics) so the existing Prometheus + Grafana stack can scrape them.
//
// Traces, metrics and logs are also sent over OTLP/gRPC to OTEL_EXPORTER_OTLP_ENDPOINT
// (the SigNoz OTel Collector, e.g. http://signoz-otel-collector.signoz.svc.cluster.local:4317).
// When that variable is unset, OTLP export is disabled and logs only go to stderr.
// The standard OTEL_* variables (OTEL_SERVICE_NAME, OTEL_RESOURCE_ATTRIBUTES,
// OTEL_TRACES_SAMPLER, ...) are honoured.
package telemetry

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.40.0"
)

const defaultMetricsAddr = ":9464"

// Init installs the global tracer, meter and logger providers, starts the
// Prometheus /metrics endpoint and, when OTLP is configured, makes slog (and the
// standard log package) write to both stderr and OTLP.
// The returned function flushes and stops the exporters.
func Init(ctx context.Context, serviceName string) (shutdown func(context.Context) error, err error) {
	var shutdowns []func(context.Context) error
	shutdown = func(ctx context.Context) error {
		var errs error
		for _, fn := range shutdowns {
			errs = errors.Join(errs, fn(ctx))
		}
		return errs
	}
	defer func() {
		if err != nil {
			err = errors.Join(err, shutdown(ctx))
		}
	}()

	otlpEnabled := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") != ""

	// Later options win, so OTEL_SERVICE_NAME / OTEL_RESOURCE_ATTRIBUTES override the name passed in.
	res, err := resource.New(ctx,
		resource.WithAttributes(semconv.ServiceName(serviceName)),
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
	)
	if err != nil {
		return nil, err
	}

	// Carry the trace context across gRPC calls (W3C traceparent header)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Metrics (gRPC/HTTP request duration, DB pool stats, Go runtime, ...).
	// The same instruments feed both Prometheus (pull) and SigNoz (OTLP push).
	// A private registry keeps client_golang's default Go and process collectors out:
	// the OTel runtime instrumentation below covers Go, and cAdvisor covers the container.
	registry := prometheus.NewRegistry()
	promExporter, err := otelprom.New(otelprom.WithRegisterer(registry))
	if err != nil {
		return nil, err
	}
	metricOpts := []sdkmetric.Option{
		sdkmetric.WithReader(promExporter),
		sdkmetric.WithResource(res),
	}
	if otlpEnabled {
		metricExporter, err := otlpmetricgrpc.New(ctx)
		if err != nil {
			return nil, err
		}
		metricOpts = append(metricOpts, sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)))
	}
	mp := sdkmetric.NewMeterProvider(metricOpts...)
	shutdowns = append(shutdowns, mp.Shutdown)
	otel.SetMeterProvider(mp)

	if err := runtime.Start(runtime.WithMeterProvider(mp)); err != nil {
		return nil, err
	}

	metricsServer, err := serveMetrics(registry)
	if err != nil {
		return nil, err
	}
	shutdowns = append(shutdowns, metricsServer.Shutdown)

	if !otlpEnabled {
		slog.Info("OTLP export disabled, OTEL_EXPORTER_OTLP_ENDPOINT is not set")
		return shutdown, nil
	}

	// Traces
	traceExporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res),
	)
	shutdowns = append(shutdowns, tp.Shutdown)
	otel.SetTracerProvider(tp)

	// Logs: keep printing to stderr for kubectl logs, and also ship them with trace IDs
	logExporter, err := otlploggrpc.New(ctx)
	if err != nil {
		return nil, err
	}
	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(logExporter)),
		sdklog.WithResource(res),
	)
	shutdowns = append(shutdowns, lp.Shutdown)
	global.SetLoggerProvider(lp)

	slog.SetDefault(slog.New(slog.NewMultiHandler(
		slog.NewTextHandler(os.Stderr, nil),
		otelslog.NewHandler(serviceName, otelslog.WithLoggerProvider(lp)),
	)))

	return shutdown, nil
}

// serveMetrics exposes the Prometheus registry on METRICS_ADDR (default :9464).
// It listens before returning so a port clash fails startup instead of being logged later.
func serveMetrics(registry *prometheus.Registry) (*http.Server, error) {
	addr := os.Getenv("METRICS_ADDR")
	if addr == "" {
		addr = defaultMetricsAddr
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	srv := &http.Server{Handler: mux, ReadHeaderTimeout: 5 * time.Second}

	go func() {
		if err := srv.Serve(lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Metrics server stopped", "err", err)
		}
	}()
	slog.Info("Serving Prometheus metrics", "addr", addr, "path", "/metrics")

	return srv, nil
}

// Flush sends buffered spans, metrics and logs, waiting at most 5 seconds. Call it last,
// after the servers have drained, so telemetry from the final requests isn't lost.
func Flush(shutdown func(context.Context) error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := shutdown(ctx); err != nil {
		slog.Error("Failed to flush telemetry", "err", err)
	}
}

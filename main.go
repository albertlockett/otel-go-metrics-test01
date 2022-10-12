package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"log"
	"net/http"
	"time"
)

func main() {

	meterName := "github.com/albertlockett/otel-go-metrics-test01"
	exporter := otelprom.New()
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter(meterName)

	go serveMetrics(exporter.Collector)

	histogram, err := meter.SyncInt64().Histogram("my-histogram")
	if err != nil {
		log.Print(err)
		panic(err)
	}

	ctx := context.Background()
	histogram.Record(ctx, 95)
	histogram.Record(ctx, 95)
	histogram.Record(ctx, 95)
	histogram.Record(ctx, 245)
	histogram.Record(ctx, 245)
	histogram.Record(ctx, 495)

	for {
		time.Sleep(30 * time.Second)
	}
}

func serveMetrics(collector prometheus.Collector) {
	registry := prometheus.NewRegistry()
	err := registry.Register(collector)
	if err != nil {
		fmt.Printf("error registering collector: %v", err)
		return
	}

	log.Printf("serving metrics at localhost:2222/metrics")
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	err = http.ListenAndServe(":2222", nil)
	if err != nil {
		fmt.Printf("error serving http: %v", err)
		return
	}
}

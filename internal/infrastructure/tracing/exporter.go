package tracing

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
)

func NewExporter(collectorUrl string) (*jaeger.Exporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(collectorUrl)))
}

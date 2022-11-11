package tracing

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"log"
)

func NewResource(ctx context.Context, serviceName, environment string) *resource.Resource {
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(semconv.ServiceNameKey.String(serviceName),
			attribute.String("environment", environment),
		),
	)
	if err != nil {
		log.Fatalf("%s: %v", "Failed to create resource", err)
	}
	return res
}

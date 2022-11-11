package tracing

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
	"vitorsavian/github-api/internal/infrastructure/env"
)

func InitProviderExporter(ctx context.Context, environ env.Environment) (func(context.Context) error, error) {
	exp, err := NewExporter(environ.CollectorUrl)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
	tp := trace.NewTracerProvider(
		trace.WithSampler(NewSampler(environ.AppEnv)),
		trace.WithBatcher(exp),
		trace.WithResource(NewResource(ctx, environ.ServiceName, environ.AppEnv)),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown, nil
}

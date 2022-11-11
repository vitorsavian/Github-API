package tracing

import (
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewSampler(env string) trace.Sampler {
	switch env {
	case "development":
		return trace.AlwaysSample()
	case "production":
		return trace.ParentBased(trace.TraceIDRatioBased(0.5))
	default:
		return trace.AlwaysSample()
	}
}

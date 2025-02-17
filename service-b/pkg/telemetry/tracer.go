package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/trace"
	"os"
)

func InitTracer() (*trace.TracerProvider, error) {
	zipkinURL := os.Getenv("OTEL_EXPORTER_ZIPKIN_ENDPOINT")
	
	exporter, err := zipkin.New(zipkinURL)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)

	return tp, nil
} 
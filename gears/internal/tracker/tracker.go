package tracker

import (
	"context"
	"os"

	"goclub/gears/internal/logger"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdk_trace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	otel_trace "go.opentelemetry.io/otel/trace"
)

var tracker otel_trace.Tracer

func InitTracker(rootCtx context.Context, trackID string, tracerURL string) {
	const fnLogName = ".InitTracker"
	defer func() {
		tracker = otel.Tracer("")
	}()

	if trackID == "" || tracerURL == "" {
		return
	}

	exporter, err := otlptracehttp.New(rootCtx, otlptracehttp.WithEndpointURL(tracerURL))
	if err != nil {
		logger.Error(rootCtx, "otel exporter error", err)
		os.Exit(-1)
	}

	rsrc, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(trackID),
			semconv.DeploymentEnvironment("route256"),
		),
	)
	if err != nil {
		logger.Error(rootCtx, "otel resource build error", err)
		os.Exit(-1)
	}

	tracerProvider := sdk_trace.NewTracerProvider(
		sdk_trace.WithBatcher(exporter),
		sdk_trace.WithResource(rsrc),
	)

	otel.SetTracerProvider(tracerProvider)

	//defer: tracker = otel.Tracer("")

}

func StartSpan(ctx context.Context, name string, opts ...otel_trace.SpanStartOption) (context.Context, otel_trace.Span) {
	ctx, span := tracker.Start(ctx, name, opts...)
	return ctx, span
}

func CloseTracker() error {
	var tracerProvider *sdk_trace.TracerProvider
	tracerProvider, _ = otel.GetTracerProvider().(*sdk_trace.TracerProvider)
	err := tracerProvider.ForceFlush(context.Background())
	if err != nil {
		return err
	}
	err = tracerProvider.Shutdown(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func GetTraceID(ctx context.Context) string {
	spanCtx := otel_trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		return spanCtx.TraceID().String()
	}
	return ""
}

func GetSpanID(ctx context.Context) string {
	spanCtx := otel_trace.SpanContextFromContext(ctx)
	if spanCtx.HasSpanID() {
		return spanCtx.SpanID().String()
	}
	return ""
}

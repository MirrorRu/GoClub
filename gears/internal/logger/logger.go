package logger

import (
	"context"
	"log/slog"
	"os"
	"runtime"

	"go.opentelemetry.io/otel/trace"
)

var log *slog.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

func GetCallerInfo(skip int) (info string) {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "func.unknown"
	}
	return runtime.FuncForPC(pc).Name()
}

func Info(ctx context.Context, msg string, args ...any) {
	richArgs := appendArgs(ctx, args...)
	log.Info(msg, richArgs...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	richArgs := appendArgs(ctx, args...)
	log.Warn(msg, richArgs...)
}

func Error(ctx context.Context, msg string, err error, args ...any) {
	args = append(args, "error", err)
	richArgs := appendArgs(ctx, args...)
	log.Error(msg, richArgs...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	richArgs := appendArgs(ctx, args...)
	log.Debug(msg, richArgs...)
}

func appendArgs(ctx context.Context, args ...any) []any {
	spanCtx := trace.SpanContextFromContext(ctx)

	if spanCtx.HasSpanID() {
		args = append(args, "span_id", spanCtx.SpanID().String())
	}

	if spanCtx.HasTraceID() {
		args = append(args, "trace_id", spanCtx.TraceID().String())
	}

	return args
}

func Logger() *slog.Logger {
	return log
}

package logger

import (
	"context"
	"log/slog"
	"os"
)

var log *slog.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

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
	//TODO extract values from context, and append to result
	return args
}

func Logger() *slog.Logger {
	return log
}

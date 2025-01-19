package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

var log *slog.Logger = slog.New(tint.NewHandler(os.Stdout, &tint.Options{
	Level:      slog.LevelDebug,
	TimeFormat: "2006-01-02 15:04:05.000",
}))

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

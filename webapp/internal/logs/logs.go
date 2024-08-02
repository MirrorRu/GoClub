package logs

import (
	"goclub/webapp/internal/error"
	"os"

	"log/slog"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

const (
	ErrCodeKey      = "code"
	ErrNameKey      = "name"
	ErrOperationKey = "operation"
	ErrSourceKey    = "source"
)

func AppError(_appErr error.AppError) {
	logger.Error(_appErr.Error(),
		ErrCodeKey, _appErr.Code(),
		ErrNameKey, _appErr.Name(),
		ErrOperationKey, _appErr.Operation(),
		ErrSourceKey, _appErr.SourceErr())

}

func Debug(_msg string, _args ...any) {
	logger.Debug(_msg, _args...)
}

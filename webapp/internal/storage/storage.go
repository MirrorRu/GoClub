package storage

import (
	"context"
	"goclub/webapp/internal/config"
	"goclub/webapp/internal/error"
	"goclub/webapp/internal/txt"
	"strings"
)

type Storage interface {
	Open() error.AppError
	Close() error.AppError
}

func NewStorage(_ctx context.Context, _cfg *config.AppConfig) (Storage, error.AppError) {
	const opName = "NewStorage"
	var (
		storage Storage
		appErr  error.AppError
	)

	switch true {
	case strings.HasPrefix(_cfg.DatabaseDSN, SQLiteDBPrefix):
		storage, appErr = NewSQLiteStorage(_cfg.DatabaseDSN)
	case IsPostgreDSN(_cfg.DatabaseDSN):
		storage, appErr = NewPostgreStorage(_cfg.DatabaseDSN)
	default:
		appErr = error.FromMessage(opName, error.ECStorageCreating, txt.ErrCantDetectStorageTypeForDSN, _cfg.DatabaseDSN)
	}
	return storage, appErr
}

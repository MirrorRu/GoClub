package storage

import (
	"goclub/webapp/internal/error"
	"strings"
)

const SQLiteDBPrefix = "sqlite:"

type SQLiteStorage struct {
	BasicStorage
}

func IsSQLiteDSN(_dsn string) bool {
	return strings.HasPrefix(_dsn, SQLiteDBPrefix)
}

func NewSQLiteStorage(dsn string) (*SQLiteStorage, error.AppError) {
	return &SQLiteStorage{}, nil
}

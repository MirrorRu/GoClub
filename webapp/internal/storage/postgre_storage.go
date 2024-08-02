package storage

import (
	"goclub/webapp/internal/error"
	"strings"
)

const PostgreDBPrefix = "postgresql:"

type PostgreStorage struct {
	BasicStorage
}

func IsPostgreDSN(_dsn string) bool {
	return strings.HasPrefix(_dsn, PostgreDBPrefix)
}

func NewPostgreStorage(_dsn string) (*PostgreStorage, error.AppError) {
	return &PostgreStorage{}, nil
}

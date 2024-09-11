package storage

import "goclub/webapp/internal/error"

type BasicStorage struct {
}

func (store *BasicStorage) Open() error.AppError {
	return nil
}

func (store *BasicStorage) Close() error.AppError {
	return nil
}

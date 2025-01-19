package logger_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"goclub/v1/common/logger"
	"testing"
)

func TestInfo(t *testing.T) {
	assert.NotPanics(t, func() {
		logger.Info(context.Background(), "info")
	})
}

func TestWarn(t *testing.T) {
	assert.NotPanics(t, func() {
		logger.Warn(context.Background(), "info")
	})
}

func TestDebug(t *testing.T) {
	assert.NotPanics(t, func() {
		logger.Debug(context.Background(), "info")
	})
}

func TestError(t *testing.T) {
	assert.NotPanics(t, func() {
		logger.Error(context.Background(), "info", errors.New("error"))
	})
}

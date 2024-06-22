package middleware

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/stretchr/testify/assert"
	"github.com/webkimru/go-keeper/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

func TestInterceptorLogger(t *testing.T) {
	t.Run("check logger implementation", func(t *testing.T) {
		// observer log message
		observedZapCore, observedLogs := observer.New(zap.InfoLevel)
		observedLogger := zap.New(observedZapCore)
		implLogger := logger.Log{Log: observedLogger.Sugar()}

		// our interceptor
		interceptor := InterceptorLogger(&implLogger)

		// check logging.Logger implementation
		switch v := interceptor.(type) {
		case logging.Logger:
			//
		default:
			t.Errorf("Expected logging.Logger, but got %T", v)
		}

		// do our custom log and check length + message
		interceptor.Log(context.Background(), 1, "test", "field")
		assert.Equal(t, 1, observedLogs.Len())
		firstLog := observedLogs.All()[0]
		assert.Equal(t, "test [field]", firstLog.Message)
	})
}

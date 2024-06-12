package middleware

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/webkimru/go-keeper/pkg/logger"
)

func InterceptorLogger(l *logger.Log) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log.Infoln(msg, fields)
	})
}

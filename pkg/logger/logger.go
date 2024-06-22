// Package logger implements logging errors.
// This package contains Zap Logger that implements Logger interface of this package.
package logger

import "go.uber.org/zap"

// Logger is the logging interface.
// Third party logger should implement it.
type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Infof(message string, args ...any)
	Infoln(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Errorf(message string, args ...any)
	Fatal(args ...any)
	Fatalf(message string, args ...any)
}

// Log contains a logger.
type Log struct {
	Log Logger
}

// NewZap implements zsp logger.
func NewZap(level string) (*Log, error) {
	// transforms to zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}
	// creates a new logger config
	cfg := zap.NewProductionConfig()
	// sets a level
	cfg.Level = lvl
	// creates s new logger with defined configuration
	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	// sets singleton
	return &Log{Log: zl.Sugar()}, nil
}

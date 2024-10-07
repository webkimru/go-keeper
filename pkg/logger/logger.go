// Package logger implements logging errors.
// This package contains Zap Logger that implements Logger interface of this package.
package logger

import (
	"go.uber.org/zap"
)

const _defaultLogLevel = "info"

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
	Log    Logger
	Level  string
	Output []string
}

// NewZap implements zsp logger.
func NewZap(opts ...Option) (*Log, error) {
	l := &Log{
		Level:  _defaultLogLevel,
		Output: []string{"stderr"}, // "path_to_logfile"
	}
	// Custom options
	for _, opt := range opts {
		opt(l)
	}
	// transforms to zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(l.Level)
	if err != nil {
		return nil, err
	}
	// creates a new logger config
	cfg := zap.NewProductionConfig()
	// sets a level
	cfg.Level = lvl
	// save to file
	cfg.OutputPaths = l.Output
	// creates s new logger with defined configuration
	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	// sets singleton
	return &Log{Log: zl.Sugar()}, nil
}

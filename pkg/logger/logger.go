package logger

import "go.uber.org/zap"

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

type Log struct {
	Log Logger
}

func NewZap(level string) (*Log, error) {
	// преобразуем текстовый уровень логирования в zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}
	// создаём новую конфигурацию логера
	cfg := zap.NewProductionConfig()
	// устанавливаем уровень
	cfg.Level = lvl
	// создаём логер на основе конфигурации
	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	// устанавливаем синглтон
	return &Log{Log: zl.Sugar()}, nil
}

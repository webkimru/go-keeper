package logger

// Option is functional options.
type Option func(log *Log)

// SetLevel is the level logging.
func SetLevel(level string) Option {
	return func(l *Log) {
		l.Level = level
	}
}

// SetOutput writes the log output to a file as well as the console/terminal, using in zap.config.
func SetOutput(output []string) Option {
	return func(l *Log) {
		l.Output = output
	}
}

package config

import (
	"go.uber.org/zap"
)

// Logger is the interface we use around zap to let us mock in tests
type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	With(...zap.Field) Logger
}

// NewLogger returns a logger based on the environment that is set
func NewLogger(v Environment) (Logger, error) {
	log, err := zap.NewProduction()
	if v == EnvDevelopment {
		log, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	return &logger{Logger: log}, nil
}

// logger wraps zap but fixes an issue where using With() would break our interface.
type logger struct {
	*zap.Logger
}

func (l *logger) With(fields ...zap.Field) Logger {
	return &logger{Logger: l.Logger.With(fields...)}
}

// Noop is a logger that does nothing.
type NoopLogger struct{}

var _ Logger = (*NoopLogger)(nil) // compile time assertion only, this uses no memory

// Info is a no-op logger
func (l *NoopLogger) Info(_ string, _ ...zap.Field) {}

// Error is a no-op logger
func (l *NoopLogger) Error(_ string, _ ...zap.Field) {}

// Debug is a no-op logger
func (l *NoopLogger) Debug(_ string, _ ...zap.Field) {}

// Fatal is a no-op logger
func (l *NoopLogger) Fatal(_ string, _ ...zap.Field) {}

// Warn is a no-op logger
func (l *NoopLogger) Warn(_ string, _ ...zap.Field) {}

// With is a no-op logger
func (l *NoopLogger) With(_ ...zap.Field) Logger {
	return l
}

// infrastructure/logger/slog_adapter.go
package logger

import (
	"inoUwU/pinu/app/domain/port"
	"log/slog"
)

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(l *slog.Logger) port.Logger {
	return &SlogLogger{logger: l}
}

func (l *SlogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *SlogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

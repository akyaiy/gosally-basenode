package logger

import "log/slog"

type appLogger interface {
	newLogger() *Logger
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
}

type Logger struct {
	appLogger
	sloglogger *slog.Logger
}

type GSLog struct {
	App  *Logger
	HTTP *Logger
	DB   *Logger
}

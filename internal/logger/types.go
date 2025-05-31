package logger

import "log/slog"

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type Log struct {
	Logger
	sloglogger *slog.Logger
}

type GSLog struct {
	App  *Log
	HTTP *Log
	DB   *Log
}

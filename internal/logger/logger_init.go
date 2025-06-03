package logger

import (
	"io"
	"log/slog"
	"os"
)

func newLogger(w io.Writer, o slog.HandlerOptions) *Log {
	return &Log{slogLogger: slog.New(slog.NewTextHandler(w, &o))}
}

func InitBaseLog() *GSLog {
	_loggers := &GSLog{
		App:  newLogger(os.Stdout, slog.HandlerOptions{Level: slog.LevelInfo}),
		HTTP: newLogger(os.Stdout, slog.HandlerOptions{Level: slog.LevelInfo}),
		DB:   newLogger(os.Stdout, slog.HandlerOptions{Level: slog.LevelInfo}),
	}
	return _loggers
}

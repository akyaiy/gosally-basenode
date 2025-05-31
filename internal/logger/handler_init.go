package logger

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

const LevelFallback = slog.Level(-8)

func createTextHandler(out io.Writer, level slog.Level, addSource bool) slog.Handler {
	return slog.NewTextHandler(out, &slog.HandlerOptions{
		Level:     level,
		AddSource: addSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				if level == LevelFallback {
					return slog.String(slog.LevelKey, "FALLBACK")
				}
			}
			if a.Key == slog.TimeKey {
				t := a.Value.Time()
				return slog.String("time", t.Format("2006-01-02 15:04:05"))
			}
			return a
		},
	})
}

func createJSONHandler(out io.Writer, level slog.Level, addSource bool) slog.Handler {
	return slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level:     level,
		AddSource: addSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				if level == LevelFallback {
					return slog.String(slog.LevelKey, "FALLBACK")
				}
			}
			return a
		},
	})
}

func InitMultiHandler(stdout bool, slogPath string, level slog.Level) *slog.Logger {
	var (
		// Whether to indicate in the logs the function, package and file from which the log came
		addSource      = level == slog.LevelDebug || level == LevelFallback
		jsonHandler    slog.Handler
		textHandler    slog.Handler
		rotatingWriter *lumberjack.Logger
	)
	if !(slogPath == "" || slogPath == "\000") {
		rotatingWriter = &lumberjack.Logger{
			Filename:   slogPath,
			MaxSize:    5,
			MaxBackups: 3,
			MaxAge:     7,
			Compress:   true,
		}
		jsonHandler = createJSONHandler(rotatingWriter, level, addSource)
	} else {
		return slog.New(NewMultiHandler(createTextHandler(os.Stdout, level, addSource)))
	}
	if stdout {
		textHandler = createTextHandler(os.Stdout, level, addSource)
		return slog.New(NewMultiHandler(textHandler, jsonHandler))
	}
	return slog.New(NewMultiHandler(jsonHandler))
}

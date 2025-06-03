package logger

func (l *Log) Info(msg string, args ...any) {
	l.slogLogger.Info(msg, args...)
}

func (l *Log) Warn(msg string, args ...any) {
	l.slogLogger.Warn(msg, args...)
}

func (l *Log) Error(msg string, args ...any) {
	l.slogLogger.Error(msg, args...)
}

func (l *Log) Debug(msg string, args ...any) {
	l.slogLogger.Debug(msg, args...)
}

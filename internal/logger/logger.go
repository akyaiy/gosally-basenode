package logger

func (l *Logger) Info(msg string, args ...any) {
	l.sloglogger.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.sloglogger.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.sloglogger.Error(msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.sloglogger.Debug(msg, args...)
}

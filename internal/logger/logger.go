package logger

func (l *Log) Info(msg string, args ...any) {
	l.sloglogger.Info(msg, args...)
}

func (l *Log) Warn(msg string, args ...any) {
	l.sloglogger.Warn(msg, args...)
}

func (l *Log) Error(msg string, args ...any) {
	l.sloglogger.Error(msg, args...)
}

func (l *Log) Debug(msg string, args ...any) {
	l.sloglogger.Debug(msg, args...)
}

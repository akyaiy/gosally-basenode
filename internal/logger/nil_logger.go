package logger

type VoidLoggerType struct{}

// Убедимся что VoidoggerType реализует интерфейс Logger
var _ Logger = (*VoidLoggerType)(nil)

var VoidLogger Logger = &VoidLoggerType{}

func (l *VoidLoggerType) Debug(args ...interface{})                   {}
func (l *VoidLoggerType) Debugf(format string, args ...interface{})   {}
func (l *VoidLoggerType) Info(args ...interface{})                    {}
func (l *VoidLoggerType) Infof(format string, args ...interface{})    {}
func (l *VoidLoggerType) Warn(args ...interface{})                    {}
func (l *VoidLoggerType) Warnf(format string, args ...interface{})    {}
func (l *VoidLoggerType) Warning(args ...interface{})                 {}
func (l *VoidLoggerType) Warningf(format string, args ...interface{}) {}
func (l *VoidLoggerType) Error(args ...interface{})                   {}
func (l *VoidLoggerType) Errorf(format string, args ...interface{})   {}
func (l *VoidLoggerType) Fatal(args ...interface{})                   {}
func (l *VoidLoggerType) Fatalf(format string, args ...interface{})   {}

func IsVoid(l *Log) bool {
	if l == nil || l.Logger == nil {
		return true
	}
	_, isVoid := l.Logger.(*VoidLoggerType)
	return isVoid
}

func IsNil(l *Log) bool {
	return l == nil || l.sloglogger == nil || l.Logger == nil || IsVoid(l)
}

func IsNotNil(l *Log) bool {
	return !IsNil(l)
}

func SetVoid(l *Log) {
	l.sloglogger = nil
	l.Logger = VoidLogger
}

// CheckAndSetVoid checks if the logger is nil or void, and sets it to VoidLogger if it is.
func CheckAndSetVoid(l *Log) {
	if IsNil(l) {
		SetVoid(l)
	}
}

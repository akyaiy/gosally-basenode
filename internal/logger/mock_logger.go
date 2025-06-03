package logger

type MockLoggerType struct{}

// Убедимся что MockLoggerType реализует интерфейс Logger
var _ Logger = (*MockLoggerType)(nil)

func (l *MockLoggerType) Debug(args ...interface{})                   {}
func (l *MockLoggerType) Debugf(format string, args ...interface{})   {}
func (l *MockLoggerType) Info(args ...interface{})                    {}
func (l *MockLoggerType) Infof(format string, args ...interface{})    {}
func (l *MockLoggerType) Warn(args ...interface{})                    {}
func (l *MockLoggerType) Warnf(format string, args ...interface{})    {}
func (l *MockLoggerType) Warning(args ...interface{})                 {}
func (l *MockLoggerType) Warningf(format string, args ...interface{}) {}
func (l *MockLoggerType) Error(args ...interface{})                   {}
func (l *MockLoggerType) Errorf(format string, args ...interface{})   {}
func (l *MockLoggerType) Fatal(args ...interface{})                   {}
func (l *MockLoggerType) Fatalf(format string, args ...interface{})   {}

func NewMockLogger() *Log {
	return &Log{
		Logger:     &MockLoggerType{},
		slogLogger: nil,
	}
}

func IsMock(l *Log) bool {
	if l == nil || l.Logger == nil {
		return true
	}
	_, isVoid := l.Logger.(*MockLoggerType)
	return isVoid
}

// IsNil checks if the logger is nil or void (MockLogger).
func IsNil(l *Log) bool {
	return l == nil || l.slogLogger == nil || l.Logger == nil || IsMock(l)
}

func IsNotNil(l *Log) bool {
	return !IsNil(l)
}

// SetVoid sets the logger to a mock logger (MockLogger).
// This is useful for testing purposes when you want to avoid actual logging.
func SetVoid(l *Log) {
	l.slogLogger = nil
	l.Logger = (*MockLoggerType)(nil)
}

// CheckAndSetVoid checks if the logger is nil or void, and sets it to MockLogger if it is.
func CheckAndSetVoid(l *Log) {
	if IsNil(l) {
		SetVoid(l)
	}
}

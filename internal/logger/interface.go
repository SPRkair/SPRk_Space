package logger

type Logger interface {
	Info(msg string, a ...interface{})
	Error(msg string, a ...interface{})
	Warning(msg string, a ...interface{})
	Debug_info(msg string, a ...interface{})
	Debug_value(msg string, a ...interface{})
	Debug_point(msg string, a ...interface{})
	Fatal(msg string, a ...interface{})
	Close()
}

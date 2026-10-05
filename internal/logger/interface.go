package logger

type Logger interface {
	Println(typeMsg string, a ...interface{})
	Printf(typeMsg string, format string, a ...interface{})
	Fatal(typeMsg string, a ...interface{})
	Close()
	DropAndTrunc() []LogMessage
}

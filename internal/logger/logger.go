package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Log struct {
	messages []LogMessage
	log      *log.Logger
	mu       sync.Mutex
}

type LogMessage struct {
	Date    int64
	TypeMsg string
	Message string
}

func NewLogger() Logger {
	logger := Log{
		messages: nil,
		log:      log.New(os.Stdout, " ", log.Ldate|log.Ltime),
	}
	return &logger
}

func (l *Log) Println(typeMsg string, a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Printf("%s: %s", typeMsg, fmt.Sprint(a...))
	l.messages = append(l.messages, LogMessage{
		Date:    time.Now().Unix(),
		TypeMsg: typeMsg,
		Message: fmt.Sprintln(a...),
	})
}

func (l *Log) Printf(typeMsg string, format string, a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.log.Printf("%s: %s", typeMsg, fmt.Sprintf(format, a...))
	l.messages = append(l.messages, LogMessage{
		Date:    time.Now().Unix(),
		TypeMsg: typeMsg,
		Message: fmt.Sprintf(format, a...),
	})
}

func (l *Log) Fatal(typeMsg string, a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.messages = append(l.messages, LogMessage{
		Date:    time.Now().Unix(),
		TypeMsg: typeMsg,
		Message: fmt.Sprintln(a...),
	})
	l.log.Fatalf("%s: %s", typeMsg, fmt.Sprint(a...))
}

func (l *Log) Close() {
	l.Println("Debug_info", "log.Close")
}

func (l *Log) DropAndTrunc() []LogMessage {
	l.mu.Lock()
	defer l.mu.Unlock()
	msgs := l.messages
	l.messages = nil
	return msgs
}

package logger

import (
	"SPRk_Space/internal/config"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type LogMsgType string

const (
	Info        LogMsgType = "Info"
	Error       LogMsgType = "Error"
	Fatal       LogMsgType = "Fatal"
	Warning     LogMsgType = "Warning"
	Debug_info  LogMsgType = "Debug_info"
	Debug_point LogMsgType = "Debug_point"
	Debug_value LogMsgType = "Debug_value"
)

type Log struct {
	config   *config.Logger
	messages []LogMessage
	log      *log.Logger
	mu       sync.Mutex
	file     *os.File
}

type LogMessage struct {
	Date    string
	TypeMsg LogMsgType
	Message string
}

func NewLogger(cfg *config.Logger) (Logger, error) {
	filename := cfg.Log_dir + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15-04-05") + cfg.Log_filename
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Logger: NewLogger: os.OpenFile: %w", err)
	}
	logger := Log{
		config:   cfg,
		messages: nil,
		log:      log.New(os.Stdout, " ", log.Ldate|log.Ltime),
		file:     f,
	}
	return &logger, nil
}

func (l *Log) writeFile(msg LogMessage) error {
	if l.config.Write_file {
		str := fmt.Sprintf("%s [%s] %s\n", msg.Date, msg.TypeMsg, msg.Message)
		_, err := l.file.WriteString(str)
		if err != nil {
			return fmt.Errorf("writeFile: %w", err)
		}
	}
	return nil
}

func (l *Log) printf(typeMsg LogMsgType, format string, a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.log.Printf("%s: %s", typeMsg, fmt.Sprintf(format, a...))

	msg := LogMessage{
		Date:    time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		TypeMsg: typeMsg,
		Message: fmt.Sprintf(format, a...),
	}
	l.writeFile(msg)
	l.messages = append(l.messages, msg)
}

func (l *Log) Fatal(msg string, a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.messages = append(l.messages, LogMessage{
		Date:    time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
		TypeMsg: Fatal,
		Message: msg + " " + fmt.Sprintln(a...),
	})
	l.log.Fatalf("Fatal: %s", fmt.Sprint(a...))
}

func (l *Log) Info(msg string, a ...interface{}) { l.printf(Info, msg, a...) }

func (l *Log) Error(msg string, a ...interface{}) { l.printf(Error, msg, a...) }

func (l *Log) Warning(msg string, a ...interface{}) { l.printf(Warning, msg, a...) }

func (l *Log) Debug_info(msg string, a ...interface{}) { l.printf(Debug_info, msg, a...) }

func (l *Log) Debug_value(msg string, a ...interface{}) { l.printf(Debug_value, msg, a...) }

func (l *Log) Debug_point(msg string, a ...interface{}) { l.printf(Debug_point, msg, a...) }

func (l *Log) Close() {
	err := l.file.Close()
	if err != nil {
		l.Error("Logger: Close: file.Close:  %v", err)
	}
	l.Debug_info("log.Close")
}

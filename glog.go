package glog

import (
	"fmt"
	"time"
)

func GenerateLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

func (l *Logger) log(level LogLevel, message string) {
	if l.level >= level {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		levelString := ""

		switch level {
		case Error:
			levelString = "ERROR"
		case Warn:
			levelString = "WARN"
		case Info:
			levelString = "INFO"
		case Debug:
			levelString = "DEBUG"
		}

		logMessage := fmt.Sprintf("[%s] [%s] %s\n", timestamp, levelString, message)

		fmt.Fprint(l.writer, logMessage)
	}
}

func (l *Logger) Error(message string) {
	l.log(Error, message)
}

func (l *Logger) Warn(message string) {
	l.log(Warn, message)
}

func (l *Logger) Info(message string) {
	l.log(Info, message)
}

func (l *Logger) Debug(message string) {
	l.log(Debug, message)
}

package welog

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func GenerateLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

func (l *Logger) log(level LogLevel, message string) {
	if l.level >= level {
		timestamp := time.Now().Format(timestampFormat)
		levelString := ""
		colorFunc := color.New(color.FgWhite).SprintFunc()

		switch level {
		case Error:
			levelString = "ERROR"
			colorFunc = color.New(color.FgRed).SprintFunc()
		case Warn:
			levelString = "WARN"
			colorFunc = color.New(color.FgYellow).SprintFunc()
		case Info:
			levelString = "INFO"
			colorFunc = color.New(color.FgGreen).SprintFunc()
		case Debug:
			levelString = "DEBUG"
			colorFunc = color.New(color.FgBlue).SprintFunc()
		}

		logMessage := fmt.Sprintf("[%s] [%s] %s\n", timestamp, colorFunc(levelString), message)

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

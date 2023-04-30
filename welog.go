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
		levelString := levelToString(level)
		colorFunc := color.New(color.FgWhite).SprintFunc()

		switch level {
		case Error:
			colorFunc = color.New(color.FgRed).SprintFunc()
		case Warn:
			colorFunc = color.New(color.FgYellow).SprintFunc()
		case Info:
			colorFunc = color.New(color.FgGreen).SprintFunc()
		case Debug:
			colorFunc = color.New(color.FgBlue).SprintFunc()
		}

		logMessage := fmt.Sprintf("[%s] [%s] %s\n", timestamp, levelString, message)

		fmt.Fprint(l.writer, colorFunc(logMessage))
	}
}

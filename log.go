package welog

import (
	"fmt"
	"time"
)

func (l *Logger) log(level LogLevel, message string) {
	if l.level >= level {
		timestamp := time.Now().Format(timestampFormat)
		levelString := levelToString(level)
		colorFunc := colorType(level)

		logMessage := fmt.Sprintf("[%s] [%s] %s\n", timestamp, levelString, message)

		fmt.Fprint(l.writer, colorFunc(logMessage))
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

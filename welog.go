package welog

import (
	"fmt"
	"time"
)

func GenerateLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

func (l *Logger) log(level LogLevel, message string) {
	if l.level >= level {
		timestamp := time.Now().Format(timestampFormat)
		levelString := levelToString(level)
		colorFunc := colorType(level)

		logMessage := fmt.Sprintf("[%s] [%s] %s\n", timestamp, levelString, message)

		fmt.Fprint(l.writer, colorFunc(logMessage))
	}
}

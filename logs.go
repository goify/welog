package welog

import (
	"fmt"
	"time"
)

func (logger *Logger) Log(level LogLevel, message string) {
	if logger.level >= level {
		timestamp := time.Now().Format(TimestampFormat)
		levelString := LevelToString(level, string(logger.mode))

		logMessage := fmt.Sprintf("[%s] [%s] %s\n", timestamp, levelString, message)

		fmt.Fprint(logger.writer, logMessage)

		if logger.logFile != nil && level == Error {
			fmt.Fprint(logger.logFile, logMessage)
		}
	}
}

func (l *Logger) Error(message string) {
	l.Log(Error, message)
}

func (l *Logger) Warn(message string) {
	l.Log(Warn, message)
}

func (l *Logger) Info(message string) {
	l.Log(Info, message)
}

func (l *Logger) Debug(message string) {
	l.Log(Debug, message)
}

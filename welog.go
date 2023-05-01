package welog

import (
	"fmt"
	"time"
)

func GenerateLogger(level LogLevel, mode LogMode) *Logger {
	return &Logger{
		level: level,
		mode:  mode,
	}
}

func (l *Logger) Log(level LogLevel, message string) {
	if l.level >= level {
		timestamp := time.Now().Format(TimestampFormat)
		levelString := LevelToString(level, string(l.mode))

		logMessage := fmt.Sprintf("[%s] [%s] %s\n", timestamp, levelString, message)

		fmt.Fprint(l.writer, logMessage)
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

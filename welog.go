package welog

import (
	"fmt"
	"io"
	"os"
	"time"
)

func GenerateLogger(level LogLevel, mode LogMode, writeFile LogWrite) *Logger {
	var writer io.Writer
	var file *os.File
	var err error

	if writeFile {
		file, err = os.OpenFile("welog-errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create log file: %v\n", err)
			writer = os.Stderr
		} else {
			writer = io.MultiWriter(os.Stderr, file)
		}
	} else {
		writer = os.Stderr
	}

	logger := &Logger{
		level:     level,
		mode:      mode,
		writer:    writer,
		writeFile: writeFile,
		logFile:   file,
	}

	return logger
}

func (l *Logger) Log(level LogLevel, message string) {
	if l.level >= level {
		timestamp := time.Now().Format(TimestampFormat)
		levelString := LevelToString(level, string(l.mode))

		logMessage := fmt.Sprintf("[%s] [%s] %s\n", timestamp, levelString, message)

		fmt.Fprint(l.writer, logMessage)

		if l.logFile != nil && level == Error {
			fmt.Fprint(l.logFile, logMessage)
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

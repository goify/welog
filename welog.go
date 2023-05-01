package welog

import (
	"fmt"
	"io"
	"os"
	"time"
)

func GenerateLogger(options ...func(*Logger)) *Logger {
	logger := &Logger{level: Info, mode: Basic}

	for _, option := range options {
		option(logger)
	}

	return logger
}

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

func (logger *Logger) SetOutput(writer io.Writer) {
	logger.writer = writer
}

func (logger *Logger) SetFormatter(formatter func(LogLevel, string) string) {
	logger.formatter = formatter
}

func WithLogLevel(level LogLevel) func(*Logger) {
	return func(logger *Logger) {
		logger.level = level
	}
}

func WithLogMode(mode LogMode) func(*Logger) {
	return func(logger *Logger) {
		logger.mode = mode
	}
}

func WithLogFile(hasLogFile LogWrite) func(*Logger) {
	return func(logger *Logger) {
		var writer io.Writer
		var file *os.File
		var err error

		if hasLogFile {
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

		logger.writeFile = hasLogFile
		logger.writer = writer
		logger.logFile = file
	}
}

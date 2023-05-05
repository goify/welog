package welog

import (
	"fmt"
	"io"
	"os"
)

func GenerateLoggerBuilder() *LoggerBuilder {
	return &LoggerBuilder{
		logger: &Logger{
			level:  Info,
			mode:   Basic,
			writer: os.Stderr,
		},
	}
}

func (builder *LoggerBuilder) WithLogLevel(level LogLevel) *LoggerBuilder {
	builder.logger.level = level
	return builder
}

func (builder *LoggerBuilder) WithLogMode(mode LogMode) *LoggerBuilder {
	builder.logger.mode = mode
	return builder
}

func (builder *LoggerBuilder) WithLogFile(hasLogFile LogWrite) *LoggerBuilder {
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

	builder.logger.writeFile = hasLogFile
	builder.logger.writer = writer
	builder.logger.logFile = file

	return builder
}

func (builder *LoggerBuilder) Build() *Logger {
	return builder.logger
}

func (logger *Logger) SetOutput(writer io.Writer) {
	logger.writer = writer
}

func (logger *Logger) SetFormatter(formatter func(LogLevel, string) string) {
	logger.formatter = formatter
}

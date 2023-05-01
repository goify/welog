package welog

import (
	"io"
	"os"
)

type LogLevel int

type LogMode string

type LogWrite bool

type Logger struct {
	level     LogLevel
	mode      LogMode
	writer    io.Writer
	writeFile LogWrite
	logFile   *os.File
	formatter func(LogLevel, string) string
}

type ColorFormatter struct{}

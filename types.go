package welog

import "io"

type LogLevel int

type LogMode string

type Logger struct {
	level     LogLevel
	mode      LogMode
	writer    io.Writer
	formatter func(LogLevel, string) string
}

type ColorFormatter struct{}

package welog

import "io"

type LogLevel int

type Logger struct {
	level     LogLevel
	writer    io.Writer
	formatter func(LogLevel, string) string
}

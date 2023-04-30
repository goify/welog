package welog

import (
	"io"

	"github.com/fatih/color"
)

func (l *Logger) SetOutput(w io.Writer) {
	l.writer = w
}

func (l *Logger) SetFormatter(formatter func(LogLevel, string) string) {
	l.formatter = formatter
}

func levelToString(level LogLevel) string {
	switch level {
	case Error:
		return "ERROR"
	case Warn:
		return "WARN"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	default:
		return ""
	}
}

func colorType(level LogLevel) func(a ...interface{}) string {
	switch level {
	case Error:
		return color.New(color.FgRed).SprintFunc()
	case Warn:
		return color.New(color.FgYellow).SprintFunc()
	case Info:
		return color.New(color.FgGreen).SprintFunc()
	case Debug:
		return color.New(color.FgBlue).SprintFunc()
	default:
		return color.New(color.FgWhite).SprintFunc()
	}
}

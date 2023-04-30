package welog

import "io"

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

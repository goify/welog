package welog

import "io"

func (l *Logger) SetOutput(w io.Writer) {
	l.writer = w
}

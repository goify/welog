package welog

func (l *Logger) Error(message string) {
	l.log(Error, message)
}

func (l *Logger) Warn(message string) {
	l.log(Warn, message)
}

func (l *Logger) Info(message string) {
	l.log(Info, message)
}

func (l *Logger) Debug(message string) {
	l.log(Debug, message)
}

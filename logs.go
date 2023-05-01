package welog

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

package welog

func GenerateLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

package welog

import "github.com/fatih/color"

func levelToString(level LogLevel, mode string) string {
	switch level {
	case Error:
		if mode == string(Basic) {
			return "ERROR"
		}
		return color.New(color.FgRed).SprintFunc()("ERROR")
	case Warn:
		if mode == string(Basic) {
			return "WARN"
		}
		return color.New(color.FgYellow).SprintFunc()("WARN")
	case Info:
		if mode == string(Basic) {
			return "INFO"
		}
		return color.New(color.FgGreen).SprintFunc()("INFO")
	case Debug:
		if mode == string(Basic) {
			return "DEBUG"
		}
		return color.New(color.FgBlue).SprintFunc()("DEBUG")
	default:
		if mode == string(Basic) {
			return "UNKNOWN"
		}
		return color.New(color.FgWhite).SprintFunc()("UNKNOWN")
	}
}

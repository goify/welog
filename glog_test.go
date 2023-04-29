package glog

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var buf bytes.Buffer

	logger := Glog(Info)

	logger.SetOutput(&buf)

	timestamp := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"

	logger.Debug("Debug message")

	infoMessage := "Info message"
	logger.Info(infoMessage)

	warnMessage := "Warn message"
	logger.Warn(warnMessage)

	errorMessage := "Error message"
	logger.Error(errorMessage)

	infoLog := timestamp + " " + "[INFO]" + " " + infoMessage
	warnLog := timestamp + " " + "[WARN]" + " " + warnMessage
	errorLog := timestamp + " " + "[ERROR]" + " " + errorMessage

	expected := infoLog + "\n" + warnLog + "\n" + errorLog + "\n"

	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Log output does not contain expected messages. Got:\n%s\nExpected:\n%s", buf.String(), expected)
	}
}

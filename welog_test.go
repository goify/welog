package welog

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/fatih/color"
)

func TestBasicMode(t *testing.T) {
	var buf bytes.Buffer

	logger := GenerateLogger(Info, "basic")

	logger.SetOutput(&buf)

	timestamp := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"

	logger.Debug("Debug message")

	infoMessage := "Info message"
	logger.Info(infoMessage)

	warnMessage := "Warn message"
	logger.Warn(warnMessage)

	errorMessage := "Error message"
	logger.Error(errorMessage)

	infoLog := timestamp + " " + "[INFO]" + " " + infoMessage + "\n"
	warnLog := timestamp + " " + "[WARN]" + " " + warnMessage + "\n"
	errorLog := timestamp + " " + "[ERROR]" + " " + errorMessage + "\n"

	expected := infoLog + warnLog + errorLog

	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Log output does not contain expected messages. Got:\n%s\nExpected:\n%s", buf.String(), expected)
	}
}

func TestColourfulMode(t *testing.T) {
	var buf bytes.Buffer

	logger := GenerateLogger(Info, "colourful")

	logger.SetOutput(&buf)

	timestamp := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"

	logger.Debug("Debug message")

	infoMessage := "Info message"
	logger.Info(infoMessage)

	warnMessage := "Warn message"
	logger.Warn(warnMessage)

	errorMessage := "Error message"
	logger.Error(errorMessage)

	infoLog := timestamp + " " + "[" + color.New(color.FgGreen).SprintFunc()("INFO") + "]" + " " + infoMessage + "\n"
	warnLog := timestamp + " " + "[" + color.New(color.FgYellow).SprintFunc()("WARN") + "]" + " " + warnMessage + "\n"
	errorLog := timestamp + " " + "[" + color.New(color.FgRed).SprintFunc()("ERROR") + "]" + " " + errorMessage + "\n"

	expected := infoLog + warnLog + errorLog

	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Log output does not contain expected messages. Got:\n%s\nExpected:\n%s", buf.String(), expected)
	}
}

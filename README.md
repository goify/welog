# Welog

![build](https://github.com/goify/welog/workflows/build/badge.svg)
![license](https://img.shields.io/github/license/goify/welog?color=success)
![Go version](https://img.shields.io/github/go-mod/go-version/goify/welog)
[![GoDoc](https://godoc.org/github.com/goify/welog?status.svg)](https://godoc.org/github.com/goify/welog)

A simple and colorful logging module for Go. It allows developers to easily log messages at different log levels, including Debug, Info, Warn, and Error.

## Install

To install the module, run the following command:

```bash
go get github.com/goify/welog
```

## Usage

To use the module in your Go program, import it using the following code:

```go
import "github.com/goify/welog"
```

Then, create a logger using the GenerateLogger function:

```go
logger := welog.GenerateLoggerBuilder().Build()
```

You can customize the logger by chaining one or more options:

```go
logger := welog.GenerateLoggerBuilder()
  .WithLogLevel(welog.Info), // By default `Info`
  .WithLogMode(welog.Colorful), // By default `Basic`
  .WithLogFile(true), // By default `false`, to give access create and write log file for errors
  .Build()
```

## Options

### WithLogLevel

Set the log level for the logger. Available levels are Debug, Info, Warn, and Error.

### WithLogMode

Set the log mode for the logger. Available modes are Basic and Colorful.

### WithLogFile

Set whether to write logs to a file in addition to the console. If true, a file called welog-errors.log will be created in the current directory and error logs will be written to it.

## Log

Once you have a logger, you can use its methods to log messages at different levels:

```go
logger.Debug("Debug message")
logger.Info("Info message")
logger.Warn("Warn message")
logger.Error("Error message")
```

By default, logs will be written to standard error in the following format:

```bash
[<timestamp>] [<level>] <message>
```

Where `<timestamp>` is the current time in the 2006-01-02 15:04:05 format, `<level>` is the log level, and `<message>` is the message to be logged.

### Colorized output

If the log mode is set to Color, the log level will be colorized in the console output:

- `Debug`: white
- `Info`: green
- `Warn`: yellow
- `Error`: red

## Customization

You can customize the output format of the logs by passing a formatting function to the `SetFormatter` method of a logger:

```go
logger.SetFormatter(func(level LogLevel, message string) string {
  return fmt.Sprintf("%s [%s] %s", level, time.Now().Format("2006-01-02"), message)
})
```

## Example

```go
package main

import "github.com/goify/welog"


func main() {
  logger := welog.GenerateLogger()

  // log some messages
  logger.Error("An error occurred")
  logger.Warn("A warning occurred")
  logger.Info("An informational message")
  logger.Debug("A debug message")
}
```

## Testing

```bash
go test
```

## Support

Welog is an MIT-licensed open source project. It can grow thanks to the sponsors and support.

## License

Welog is [MIT licensed](LICENSE).

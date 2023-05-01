# Welog

![build](https://github.com/iamando/welog/workflows/build/badge.svg)
![license](https://img.shields.io/github/license/iamando/welog?color=success)
![Go version](https://img.shields.io/github/go-mod/go-version/iamando/welog)
[![GoDoc](https://godoc.org/github.com/iamando/welog?status.svg)](https://godoc.org/github.com/iamando/welog)

A simple and colorful logging module for Go. It allows developers to easily log messages at different log levels, including Debug, Info, Warn, and Error.

## Install

To install the module, run the following command:

```bash
go get github.com/iamando/welog
```

## Usage

To use the module in your Go program, import it using the following code:

```go
import "github.com/iamando/welog"
```

Then, create a logger using the GenerateLogger function:

```go
logger := welog.GenerateLogger()
```

You can customize the logger by passing it one or more options:

```go
logger := welog.GenerateLogger(
    welog.WithLogLevel(welog.Info), // By default `Info`
    welog.WithLogMode(welog.Colorful), // By default `Basic`
    welog.WithLogFile(true), // By default `false`, to give access create and write log file for errors
)
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

## Example

```go
package main

import "github.com/iamando/welog"


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

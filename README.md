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

The **GenerateLogger** function takes a **LogLevel** parameter, which specifies the minimum log level to output. For example, if you set the log level to **Info**, the logger will output all log messages at the **Info**, **Warn**, and **Error** levels.

Once you have a logger, you can use its methods to log messages at different levels:

```go
logger.Debug("Debug message")
logger.Info("Info message")
logger.Warn("Warn message")
logger.Error("Error message")
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

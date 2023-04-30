<p align="center">
  <img src="./docs/logo.svg">
</p>

![build](https://github.com/iamando/welog/workflows/build/badge.svg)
![license](https://img.shields.io/github/license/iamando/welog?color=success)
![Go version](https://img.shields.io/github/go-mod/go-version/iamando/welog)
[![GoDoc](https://godoc.org/github.com/iamando/welog?status.svg)](https://godoc.org/github.com/iamando/welog)

A simple and colourful logger module for golang with timestamp support.

## Install

```bash
# v1 - minified logger module
go get -u github.com/iamando/welog/v1

# v2 - minified and colourful logger module
go get -u github.com/iamando/welog/v2
```

## Usage

```go
package main


import (
  welog "github.com/iamando/welog"  // use `/v1` or `/v2` if you need specific version
)


func ExampleUsage() {
 // create a new logger with info level
 logger := welog.GenerateLogger(Info)

 // log some messages
 logger.Error("An error occurred")
 logger.Warn("A warning occurred")
 logger.Info("An informational message")
 logger.Debug("A debug message")

 // set the logger level to debug
 logger.level = Debug

 // log some more messages
 logger.Error("Another error occurred")
 logger.Warn("Another warning occurred")
 logger.Info("Another informational message")
 logger.Debug("Another debug message")
}

func main() {
 ExampleUsage()
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

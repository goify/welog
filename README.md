# Glog

![build](https://github.com/iamando/glog/workflows/build/badge.svg)
![license](https://img.shields.io/github/license/iamando/glog?color=success)
![Go version](https://img.shields.io/github/go-mod/go-version/iamando/glog)

A simple and colourful logger module for golang with timestamp support.

## Install

```bash
# v1 - minified logger module
go get -u github.com/iamando/glog/v1
```

## Usage

```go
package main


import (
  glog "github.com/iamando/glog"  // use `/v1` if you need specific version
)


func ExampleUsage() {
 // create a new logger with info level
 logger := glog.GenerateLogger(Info)

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

Glog is an MIT-licensed open source project. It can grow thanks to the sponsors and support.

## License

Glog is [MIT licensed](LICENSE).

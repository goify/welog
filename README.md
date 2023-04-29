# Glog

![build](https://github.com/iamando/glog/workflows/build/badge.svg)
![license](https://img.shields.io/github/license/iamando/glog?color=success)
![Go version](https://img.shields.io/github/go-mod/go-version/iamando/glog)

A simple logger module for golang with timestamp support.

## Install

```bash
go get -u github.com/iamando/glog
```

## Usage

```go
package main


import "github.com/iamando/glog"


func ExampleUsage() {
 // create a new logger with info level
 logger := Glog(Info)

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

## Support

glog is an MIT-licensed open source project. It can grow thanks to the sponsors and support.

## License

glog is [MIT licensed](LICENSE).

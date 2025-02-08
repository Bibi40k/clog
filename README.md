# GLog

GLog is a simple logging library for Go that supports different log levels and colorized output in the terminal.

## Installation

To install the library, run:

```sh
go get github.com/Bibi40k/glog/glog
```

## Log Levels

The following log levels are supported:

- INFO
- DEBUG
- ERROR
- WARNING
- FATAL
- SUCCESS

Each log level has a corresponding color when displayed in the terminal.

## Usage

To use the Color Logger in your project, import the package and create an instance of the `Logger` struct. You can then use the `Info`, `Warning`, and `Error` methods to log messages.

### Example

```go
package main

import (
    "github.com/Bibi40k/glog"
)

func main() {
    glog.Info("This is an informational message")
    glog.Debug("This is a debug message")
    glog.Error("This is an error message")
    glog.Warning("This is a warning message")
    glog.Fatal("This is a fatal message")
    glog.Success("This is a success message")

    glog.InfoAndSave("This is an informational message", "log/info.log")
    glog.DebugAndSave("This is a debug message", "log/debug.log")
    glog.ErrorAndSave("This is an error message", "log/error.log")
    glog.WarningAndSave("This is a warning message", "log/warning.log")
    glog.FatalAndSave("This is a fatal message", "log/fatal.log")
    glog.SuccessAndSave("This is a success message", "log/success.log")
}
```

### Logging Messages

To log a message with a specific level, use the corresponding function:

```go
glog.Info("This is an informational message")
glog.Debug("This is a debug message")
glog.Error("This is an error message")
glog.Warning("This is a warning message")
glog.Fatal("This is a fatal message")
glog.Success("This is a success message")
```

## Methods

- **Info(message string)**: Logs an informational message in green.
- **Warning(message string)**: Logs a warning message in yellow.
- **Error(message string)**: Logs an error message in red.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
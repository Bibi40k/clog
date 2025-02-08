# Glog

Glog este un pachet de logging pentru Go care permite logarea mesajelor cu diferite niveluri și salvarea acestora într-un fișier.

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
    "glog"
)

func main() {
    glog.Info("This is an info message")
    glog.Infof("This is a formatted info message: %s", "example")
    glog.InfoAndSave("This is an info message saved to a file", "log.txt")

    glog.Debug("This is a debug message")
    glog.Debugf("This is a formatted debug message: %s", "example")
    glog.DebugAndSave("This is a debug message saved to a file", "log.txt")

    glog.Error("This is an error message")
    glog.Errorf("This is a formatted error message: %s", "example")
    glog.ErrorAndSave("This is an error message saved to a file", "log.txt")

    glog.Warning("This is a warning message")
    glog.Warningf("This is a formatted warning message: %s", "example")
    glog.WarningAndSave("This is a warning message saved to a file", "log.txt")

    glog.Fatal("This is a fatal message")
    glog.Fatalf("This is a formatted fatal message: %s", "example")
    glog.FatalAndSave("This is a fatal message saved to a file", "log.txt")

    glog.Success("This is a success message")
    glog.Successf("This is a formatted success message: %s", "example")
    glog.SuccessAndSave("This is a success message saved to a file", "log.txt")
}
```

### Logging Messages

To log a message with a specific level, use the corresponding function:

```go
glog.Info("This is an informational message")
glog.Infof("This is a formatted info message: %s", "example")
glog.InfoAndSave("This is an info message saved to a file", "log.txt")

glog.Debug("This is a debug message")
glog.Debugf("This is a formatted debug message: %s", "example")
glog.DebugAndSave("This is a debug message saved to a file", "log.txt")

glog.Error("This is an error message")
glog.Errorf("This is a formatted error message: %s", "example")
glog.ErrorAndSave("This is an error message saved to a file", "log.txt")

glog.Warning("This is a warning message")
glog.Warningf("This is a formatted warning message: %s", "example")
glog.WarningAndSave("This is a warning message saved to a file", "log.txt")

glog.Fatal("This is a fatal message")
glog.Fatalf("This is a formatted fatal message: %s", "example")
glog.FatalAndSave("This is a fatal message saved to a file", "log.txt")

glog.Success("This is a success message")
glog.Successf("This is a formatted success message: %s", "example")
glog.SuccessAndSave("This is a success message saved to a file", "log.txt")
```

## Methods

- **Info(message string)**: Logs an informational message in green.
- **Infof(format string, args ...interface{})**: Logs a formatted informational message in green.
- **InfoAndSave(message, filepath string)**: Logs an informational message in green and saves it to a file.
- **Warning(message string)**: Logs a warning message in yellow.
- **Warningf(format string, args ...interface{})**: Logs a formatted warning message in yellow.
- **WarningAndSave(message, filepath string)**: Logs a warning message in yellow and saves it to a file.
- **Error(message string)**: Logs an error message in red.
- **Errorf(format string, args ...interface{})**: Logs a formatted error message in red.
- **ErrorAndSave(message, filepath string)**: Logs an error message in red and saves it to a file.
- **Debug(message string)**: Logs a debug message in blue.
- **Debugf(format string, args ...interface{})**: Logs a formatted debug message in blue.
- **DebugAndSave(message, filepath string)**: Logs a debug message in blue and saves it to a file.
- **Fatal(message string)**: Logs a fatal message in magenta.
- **Fatalf(format string, args ...interface{})**: Logs a formatted fatal message in magenta.
- **FatalAndSave(message, filepath string)**: Logs a fatal message in magenta and saves it to a file.
- **Success(message string)**: Logs a success message in green.
- **Successf(format string, args ...interface{})**: Logs a formatted success message in green.
- **SuccessAndSave(message, filepath string)**: Logs a success message in green and saves it to a file.

## Funcții disponibile

### Log

Logează un mesaj cu un anumit nivel.

```go
func Log(level, message string)
```

### LogAndSave

Logează un mesaj cu un anumit nivel și salvează-l într-un fișier.

```go
func LogAndSave(level, message, filepath string)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
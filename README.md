# Color Logger

A simple Go library for logging messages with color-coded output.

## Installation

To install the library, run:

```sh
go get github.com/Bibi40k/gologger
```

## Usage

To use the Color Logger in your project, import the package and create an instance of the `Logger` struct. You can then use the `Info`, `Warning`, and `Error` methods to log messages.

### Example

```go
package main

import (
    "github.com/Bibi40k/gologger"
)

func main() {
    gologger.Info("This is an info message.")
    gologger.Warning("This is a warning message.")
    gologger.Error("This is an error message.")
    gologger.Debug("This is a debug message.")
    gologger.Fatal("This is a fatal message.")
    gologger.Success("This is a success message.")

    gologger.InfoAndSave("This is an info message saved to a file.", "log/info.log")
    gologger.WarningAndSave("This is a warning message saved to a file.", "log/warning.log")
    gologger.ErrorAndSave("This is an error message saved to a file.", "log/error.log")
    gologger.DebugAndSave("This is a debug message saved to a file.", "log/debug.log")
    gologger.FatalAndSave("This is a fatal message saved to a file.", "log/fatal.log")
    gologger.SuccessAndSave("This is a success message saved to a file.", "log/success.log")
}
```

## Methods

- **Info(message string)**: Logs an informational message in green.
- **Warning(message string)**: Logs a warning message in yellow.
- **Error(message string)**: Logs an error message in red.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
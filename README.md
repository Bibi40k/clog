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
    logger := gologger.NewLogger()

    logger.Info("This is an info message.")
    logger.Warning("This is a warning message.")
    logger.Error("This is an error message.")
}
```

## Methods

- **Info(message string)**: Logs an informational message in green.
- **Warning(message string)**: Logs a warning message in yellow.
- **Error(message string)**: Logs an error message in red.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
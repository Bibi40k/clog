package main

import "github.com/Bibi40k/clog/clog"

func main() {
	clog.Info("This is an info message.")
	clog.Infof("This is a formatted info message: %s", "example")
	clog.Warning("This is a warning message.")
	clog.Warningf("This is a formatted warning message: %s", "example")
	clog.Error("This is an error message.")
	clog.Errorf("This is a formatted error message: %s", "example")
	clog.Debug("This is a debug message.")
	clog.Debugf("This is a formatted debug message: %s", "example")
	// clog.Fatal("This is a fatal message.")
	// clog.Fatalf("This is a formatted fatal message: %s", "example")
	clog.Success("This is a success message.")
	clog.Successf("This is a formatted success message: %s", "example")

	clog.InfoAndSave("This is an info message saved to a file.", "log/info.log")
	clog.WarningAndSave("This is a warning message saved to a file.", "log/warning.log")
	clog.ErrorAndSave("This is an error message saved to a file.", "log/error.log")
	clog.DebugAndSave("This is a debug message saved to a file.", "log/debug.log")
	// clog.FatalAndSave("This is a fatal message saved to a file.", "log/fatal.log")
	clog.SuccessAndSave("This is a success message saved to a file.", "log/success.log")
}

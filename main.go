package main

import "github.com/Bibi40k/glog/glog"

func main() {
	glog.Info("This is an info message.")
	glog.Infof("This is a formatted info message: %s", "example")
	glog.Warning("This is a warning message.")
	glog.Warningf("This is a formatted warning message: %s", "example")
	glog.Error("This is an error message.")
	glog.Errorf("This is a formatted error message: %s", "example")
	glog.Debug("This is a debug message.")
	glog.Debugf("This is a formatted debug message: %s", "example")
	// glog.Fatal("This is a fatal message.")
	// glog.Fatalf("This is a formatted fatal message: %s", "example")
	glog.Success("This is a success message.")
	glog.Successf("This is a formatted success message: %s", "example")

	glog.InfoAndSave("This is an info message saved to a file.", "log/info.log")
	glog.WarningAndSave("This is a warning message saved to a file.", "log/warning.log")
	glog.ErrorAndSave("This is an error message saved to a file.", "log/error.log")
	glog.DebugAndSave("This is a debug message saved to a file.", "log/debug.log")
	// glog.FatalAndSave("This is a fatal message saved to a file.", "log/fatal.log")
	glog.SuccessAndSave("This is a success message saved to a file.", "log/success.log")
}

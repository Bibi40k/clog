package main

import "github.com/Bibi40k/glog/glog"

func main() {
	glog.Info("This is an info message.")
	glog.Warning("This is a warning message.")
	glog.Error("This is an error message.")
	glog.Debug("This is a debug message.")
	// glog.Fatal("This is a fatal message.")
	glog.Success("This is a success message.")

	glog.InfoAndSave("This is an info message saved to a file.", "log/info.log")
	glog.WarningAndSave("This is a warning message saved to a file.", "log/warning.log")
	glog.ErrorAndSave("This is an error message saved to a file.", "log/error.log")
	glog.DebugAndSave("This is a debug message saved to a file.", "log/debug.log")
	// glog.FatalAndSave("This is a fatal message saved to a file.", "log/fatal.log")
	glog.SuccessAndSave("This is a success message saved to a file.", "log/success.log")
}

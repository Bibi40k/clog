package glog

import (
	"fmt"
	"log"
	"os"
	fp "path/filepath"
	"runtime"
)

// Log levels and their corresponding colors
var logLevels = map[string]string{
	"INFO":    "Blue",
	"DEBUG":   "Cyan",
	"ERROR":   "Red",
	"WARNING": "Yellow",
	"FATAL":   "Magenta",
	"SUCCESS": "Green",
}

// Log a message with a given level
func Log(level, message string) {
	logMessage(level, message, getColor(level))
}

// Log a message with a given level and save it to a file
func LogAndSave(level, message, filepath string) {
	logAndSave(level, message, filepath, getColor(level))
}

// Info log
func Info(message string) {
	Log("INFO", message)
}

// Infof log with formatted message
func Infof(format string, args ...interface{}) {
	Log("INFO", fmt.Sprintf(format, args...))
}

// InfoAndSave logs an informational message and saves it to a file
func InfoAndSave(message, filepath string) {
	LogAndSave("INFO", message, filepath)
}

// Debug log
func Debug(message string) {
	Log("DEBUG", message)
}

// Debugf log with formatted message
func Debugf(format string, args ...interface{}) {
	Log("DEBUG", fmt.Sprintf(format, args...))
}

// DebugAndSave logs a debug message and saves it to a file
func DebugAndSave(message, filepath string) {
	LogAndSave("DEBUG", message, filepath)
}

// Error log
func Error(message string) {
	Log("ERROR", formatMessageWithCaller(message))
}

// Errorf log with formatted message
func Errorf(format string, args ...interface{}) {
	Log("ERROR", formatMessageWithCaller(fmt.Sprintf(format, args...)))
}

// ErrorAndSave logs an error message and saves it to a file
func ErrorAndSave(message, filepath string) {
	LogAndSave("ERROR", formatMessageWithCaller(message), filepath)
}

// Warning log
func Warning(message string) {
	Log("WARNING", message)
}

// Warningf log with formatted message
func Warningf(format string, args ...interface{}) {
	Log("WARNING", fmt.Sprintf(format, args...))
}

// WarningAndSave logs a warning message and saves it to a file
func WarningAndSave(message, filepath string) {
	LogAndSave("WARNING", message, filepath)
}

// Fatal log
func Fatal(message string) {
	Log("FATAL", formatMessageWithCaller(message))
}

// Fatalf log with formatted message
func Fatalf(format string, args ...interface{}) {
	Log("FATAL", formatMessageWithCaller(fmt.Sprintf(format, args...)))
}

// FatalAndSave logs a fatal message and saves it to a file
func FatalAndSave(message, filepath string) {
	LogAndSave("FATAL", formatMessageWithCaller(message), filepath)
}

// Success log
func Success(message string) {
	Log("SUCCESS", message)
}

// Successf log with formatted message
func Successf(format string, args ...interface{}) {
	Log("SUCCESS", fmt.Sprintf(format, args...))
}

// SuccessAndSave logs a success message and saves it to a file
func SuccessAndSave(message, filepath string) {
	LogAndSave("SUCCESS", message, filepath)
}

// logMessage logs a message with a given level
func logMessage(level, message, color string) {
	log.Printf("[%s] %s", level, colorize(message, color))
}

// logAndSave logs a message with a given level and saves it to a file
func logAndSave(level, message, filepath, color string) {
	if filepath == "" {
		log.Fatalf("Filepath must be specified")
	}

	// Extract the directory from the filepath
	dir := fp.Dir(filepath)

	// Ensure the log directory exists
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Open the log file
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Create a logger for the terminal
	terminalLogger := log.New(os.Stdout, "", log.LstdFlags)
	// Create a logger for the file
	fileLogger := log.New(file, "", log.LstdFlags)

	// Log message with color in terminal and without color in file
	terminalLogger.Printf("[%s] %s", level, colorize(message, color))
	fileLogger.Printf("[%s] %s", level, message)
}

// getColor returns the color for a given log level
func getColor(level string) string {
	if color, exists := logLevels[level]; exists {
		return color
	}
	return "Reset"
}

// formatMessageWithCaller formats a message with the caller's file and line number
func formatMessageWithCaller(message string) string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d %s", file, line, message)
}

// colorize adds color to log messages based on the level
func colorize(message, color string) string {
	var colorStart string
	switch color {
	case "Blue":
		colorStart = "\033[34m"
	case "Cyan":
		colorStart = "\033[36m"
	case "Red":
		colorStart = "\033[31m"
	case "Yellow":
		colorStart = "\033[33m"
	case "Magenta":
		colorStart = "\033[35m"
	case "Green":
		colorStart = "\033[32m"
	default:
		colorStart = "\033[0m" // Reset
	}
	colorEnd := "\033[0m" // Reset
	return fmt.Sprintf("%s%s%s", colorStart, message, colorEnd)
}

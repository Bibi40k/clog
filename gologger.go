package gologger

import (
	"fmt"
	"io"
	"log"
	"os"
	fp "path/filepath"
	"runtime"
	"time"
)

// Info log
func Info(message string) {
	logMessage("INFO", message, "Blue")
}

// InfoAndSave logs an informational message and saves it to a file
func InfoAndSave(message, filepath string) {
	logAndSave("INFO", message, filepath, "Blue")
}

// Debug log
func Debug(message string) {
	logMessage("DEBUG", message, "Cyan")
}

// DebugAndSave logs a debug message and saves it to a file
func DebugAndSave(message, filepath string) {
	logAndSave("DEBUG", message, filepath, "Cyan")
}

// Error log
func Error(message string) {
	_, file, line, _ := runtime.Caller(1)
	logMessage("ERROR", fmt.Sprintf("%s:%d %s", file, line, message), "Red")
}

// ErrorAndSave logs an error message and saves it to a file
func ErrorAndSave(message, filepath string) {
	_, file, line, _ := runtime.Caller(1)
	logAndSave("ERROR", fmt.Sprintf("%s:%d %s", file, line, message), filepath, "Red")
}

// Warning log
func Warning(message string) {
	logMessage("WARNING", message, "Yellow")
}

// WarningAndSave logs a warning message and saves it to a file
func WarningAndSave(message, filepath string) {
	logAndSave("WARNING", message, filepath, "Yellow")
}

// Fatal log
func Fatal(message string) {
	_, file, line, _ := runtime.Caller(1)
	log.Fatalf("[%s] %s:%d %s", "FATAL", file, line, colorize(message, "Magenta"))
}

// FatalAndSave logs a fatal message and saves it to a file
func FatalAndSave(message, filepath string) {
	_, file, line, _ := runtime.Caller(1)
	logAndSave("FATAL", fmt.Sprintf("%s:%d %s", file, line, message), filepath, "Magenta")
}

// Success log
func Success(message string) {
	logMessage("SUCCESS", message, "Green")
}

// SuccessAndSave logs a success message and saves it to a file
func SuccessAndSave(message, filepath string) {
	logAndSave("SUCCESS", message, filepath, "Green")
}

// logMessage logs a message with a given level
func logMessage(level, message, color string) {
	currentTime := time.Now().Format(time.RFC3339)
	log.Printf("[%s] %s %s", level, currentTime, colorize(message, color))
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

	// Create a logger that writes to both the file and the terminal
	multiWriter := io.MultiWriter(os.Stdout, file)
	logger := log.New(multiWriter, "", log.LstdFlags)

	currentTime := time.Now().Format(time.RFC3339)
	logger.Printf("[%s] %s %s", level, currentTime, colorize(message, color))
}

// colorize adds color to log messages based on the level
func colorize(message, level string) string {
	var colorStart, colorEnd string
	switch level {
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
	colorEnd = "\033[0m" // Reset
	return fmt.Sprintf("%s%s%s", colorStart, message, colorEnd)
}

package gologger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

type LogEntry struct {
	ID      int       // Unique ID of the log
	Level   string    // Log level (e.g., INFO, DEBUG, ERROR)
	Message string    // Log message
	Time    time.Time // Time of logging
}

// Info log
func Info(message string) {
	log.Printf("[%s] %s", "INFO", colorize(message, "Blue"))
}

// Debug log
func Debug(message string) {
	log.Printf("[%s] %s", "DEBUG", colorize(message, "Cyan"))
}

// Error log
func Error(message string) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("[%s] %s:%d %s", "ERROR", file, line, colorize(message, "Red"))
}

// Warning log
func Warning(message string) {
	log.Printf("[%s] %s", "WARNING", colorize(message, "Yellow"))
}

// Fatal log
func Fatal(message string) {
	_, file, line, _ := runtime.Caller(1)
	log.Fatalf("[%s] %s:%d %s", "FATAL", file, line, colorize(message, "Magenta"))
}

// Success log
func Success(message string) {
	log.Printf("[%s] %s", "SUCCESS", colorize(message, "Green"))
}

// InfoAndSave logs an informational message and saves it to a file
func InfoAndSave(message, filepath string) {
	if filepath == "" {
		Fatal("Filepath cannot be empty")
	}

	// Ensure the log directory exists
	err := os.MkdirAll("log", os.ModePerm)
	if err != nil {
		Fatal(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	// Open the log file
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Fatal(fmt.Sprintf("Failed to open log file: %v", err))
	}
	defer file.Close()

	// Create a logger that writes to both the file and the terminal
	multiWriter := io.MultiWriter(os.Stdout, file)
	logger := log.New(multiWriter, "", log.LstdFlags)

	logger.Printf("[%s] %s", "INFO", colorize(message, "Blue"))
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

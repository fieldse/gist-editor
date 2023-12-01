package logger

import (
	"fmt"
	"log"
	"os"
)

// Log info message to console
func Info(msg string, a ...interface{}) {
	log.Printf("[info] %s", fmt.Sprintf(msg, a...))
}

// Log debug message to console
func Debug(msg string, a ...interface{}) {
	log.Printf("[debug] %s", fmt.Sprintf(msg, a...))
}

// Log warn message to console
func Warn(msg string, a ...interface{}) {
	log.Printf("[warn] %s", fmt.Sprintf(msg, a...))
}

// Log error message to console
func Error(msg string, err error) {
	log.Printf("[error] %s - %s", msg, err.Error())
}

// Log fatal error to console, and exit program
func Fatal(msg string, err error) {
	log.Printf("[error -- fatal] %s - %s", msg, err.Error())
	os.Exit(1)
}

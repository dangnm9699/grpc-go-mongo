package logger

import (
	"log"
	"os"
)

var infoLogger *log.Logger
var debugLogger *log.Logger
var errorLogger *log.Logger

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Set used to initialize all logger
func Set() {}

// Info used to get info logger
func Info() *log.Logger {
	return infoLogger
}

// Debug used to get info logger
func Debug() *log.Logger {
	return debugLogger
}

// Error used to get info logger
func Error() *log.Logger {
	return errorLogger
}

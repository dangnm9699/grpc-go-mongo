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

func Set() {}

func Info() *log.Logger {
	return infoLogger
}

func Debug() *log.Logger {
	return debugLogger
}

func Error() *log.Logger {
	return errorLogger
}

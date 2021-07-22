package main

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

var logger *Logger

func (l *Logger) initialize() {
	file, err := os.OpenFile("gRPC_server_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	checkErr(err, "cannot create logging file")
	logger.infoLogger = log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.errorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

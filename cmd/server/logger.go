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

func (l *Logger) Initialize() {
	file, err := os.OpenFile("gRPC_server_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.SetPrefix("[ERR] ")
		log.Fatalln("cannot create logging file |||||", err)
	}
	logger.infoLogger = log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.errorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *Logger) Info(msg string) {
	l.infoLogger.Println(msg)
}

func (l *Logger) Error(err error, msg string) {
	l.errorLogger.Println(msg, "|||||", err)
}

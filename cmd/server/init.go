package main

import (
	"context"
	"os"
)

func init() {
	ctx = context.Background()
	logger = &Logger{}
	config = &Config{}
	mongodb = &MongoDb{}
	upsert = true
}

func checkErr(err error, msg string) {
	if err != nil {
		logger.Error(err, msg)
		os.Exit(1)
	}
}

func getEnv(key, alt string) string {
	env := os.Getenv(key)
	if len(env) > 0 {
		return env
	}
	return alt
}

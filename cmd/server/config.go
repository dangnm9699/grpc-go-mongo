package main

import (
	"github.com/joho/godotenv"
)

type Config struct {
	port string
	uri  string
	db   string
	coll string
}

var config *Config

func (cfg *Config) Initialize() {
	checkErr(godotenv.Load(), "cannot read .env file")
	config = &Config{
		port: getEnv("GRPC_EXAMPLE_PORT", "5000"),
		uri: getEnv("GRPC_EXAMPLE_URI", "mongodb://localhost:27017"),
		db: getEnv("GRPC_EXAMPLE_DATABASE", "grpc"),
		coll: getEnv("GRPC_EXAMPLE_COLLECTION", "example"),
	}
}

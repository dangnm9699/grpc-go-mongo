package main

import "context"

var ctx context.Context

func main() {
	logger.Initialize()
	config.Initialize()
	mongodb.Initialize()
	movieService := &MovieSvc{}
	checkErr(movieService.RunGRPC(ctx), "cannot starting gRPC server")
}

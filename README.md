# grpc-example

[![Go Report Card](https://goreportcard.com/badge/github.com/dangnm9699/grpc-go-mongo)](https://goreportcard.com/report/github.com/dangnm9699/grpc-go-mongo)

Create a gRPC example in Go using MongoDB as storage

* _Make sure your mongodb running_

* Download/Clean dependencies:
    ```shell
    go mod tidy
    ```
* Run server:
    ```shell
    ./bin/grpc-go-mongo.exe server
    ```
* Run client:
    ```shell
    ./bin/grpc-go-mongo.exe client
    ```
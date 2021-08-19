# grpc-example

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
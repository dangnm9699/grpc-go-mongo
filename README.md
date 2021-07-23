# grpc-example

Create a gRPC example using Go, MongoDB

* _Make sure your mongodb running_

* Create ```.env```:
    ```shell
    cp .env.example .env
    ```
* Edit your ```.env```
* Run server:
    ```shell
    go run github.com/dangnm9699/grpc-example/cmd/server
    ```
* Run client:
    ```shell
    go run github.com/dangnm9699/grpc-example/cmd/client
    ```
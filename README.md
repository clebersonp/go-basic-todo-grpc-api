# go-basic-todo-grpc-api

## .proto file
Create a file .proto to specify the api messages and services

## protoc installation
[grpc doc](https://grpc.io/docs/protoc-installation/)
```shell
apt install -y protobuf-compiler
```

## Go plugins
[Install plugin for Go](https://grpc.io/docs/languages/go/quickstart/)

## Execute Makefile
```shell
make generate-protoc
```

## Resolve dependencies
```shell
go mod tidy
```
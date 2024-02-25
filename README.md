# TODO api

## Description

A basic poc of Todo api written in Golang using grpc framework

## The proto file

Create a file with extension *.proto* to specify the api messages and services.
See [proto3 doc](https://protobuf.dev/programming-guides/proto3/).

## protoc installation

See [grpc doc](https://grpc.io/docs/protoc-installation/).

```shell
apt install -y protobuf-compiler
```

## Go plugins

See protocol compiler plugin installation for [Go](https://grpc.io/docs/languages/go/quickstart/).

## Regenerate gRPC code

See doc for [Go](https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code).

Run Makefile target to facilitate the generation of *gRPC code* after each modification to the *.proto* file

```shell
make generate-todo-protoc
```

## Resolve dependencies

Run the following Go command to synchronizes dependencies if necessary

```shell
go mod tidy
```

## Test

If using postman, we need to import *.proto* file as API.
See [doc](https://learning.postman.com/docs/sending-requests/grpc/using-service-definition/#importing-a-proto-file).

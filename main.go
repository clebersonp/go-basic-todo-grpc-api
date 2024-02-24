package main

import (
	"flag"
	"fmt"
	"github.com/clebersonp/go-basic-todo-grpc-api/middleware"
	"github.com/clebersonp/go-basic-todo-grpc-api/model"
	"github.com/clebersonp/go-basic-todo-grpc-api/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port    = flag.Int("port", 8089, "Api port number")
	network = flag.String("network", "tcp", "Kind of networks")
)

func main() {
	flag.Parse()
	listen, err := net.Listen(*network, fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("cannot create a tcp listener: %q", err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(middleware.RequestID, middleware.EnsureValidToken),
	}

	serverRegistrar := grpc.NewServer(opts...)
	proto.RegisterTaskerServer(serverRegistrar, &model.TaskerServer{})

	log.Println("Listen at:", listen.Addr())
	err = serverRegistrar.Serve(listen)
	if err != nil {
		log.Fatalf("Impossible to start up the server: %q", err)
	}
}

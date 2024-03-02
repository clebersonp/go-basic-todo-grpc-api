package main

import (
	"flag"
	"fmt"
	"github.com/clebersonp/go-basic-todo-grpc-api/middleware"
	"github.com/clebersonp/go-basic-todo-grpc-api/model/todo"
	"github.com/clebersonp/go-basic-todo-grpc-api/model/user"
	todopb "github.com/clebersonp/go-basic-todo-grpc-api/proto/todo"
	userpb "github.com/clebersonp/go-basic-todo-grpc-api/proto/user"
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
	todopb.RegisterTaskerServer(serverRegistrar, &todo.TaskerServer{})
	userpb.RegisterUsersServer(serverRegistrar, &user.ServerUser{})

	log.Println("Listen at:", listen.Addr())
	err = serverRegistrar.Serve(listen)
	if err != nil {
		log.Fatalf("Impossible to start up the server: %q", err)
	}
}

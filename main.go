package main

import (
	"context"
	"fmt"
	"github.com/clebersonp/go-basic-todo-grpc-api/model"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strings"
)

var (
	errMissingMetadata = status.Error(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Error(codes.Unauthenticated, "invalid token")
)

type taskerServer struct {
	model.UnimplementedTaskerServer
}

func (t *taskerServer) Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error) {
	log.Println("service:", t)
	log.Println("UnimplementedTaskerServer:", t.UnimplementedTaskerServer)
	log.Println("context:", ctx)

	// get metadata from request (like headers)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Println("MD:", md)
		log.Println("Authorization:", md["authorization"])
	}
	log.Println("req:", req)
	log.Println("Todo", req.GetTodo())

	// sending metadata to client
	reqId := uuid.New().String()
	header := metadata.Pairs("x-request-id", reqId)
	err := grpc.SendHeader(ctx, header)
	if err != nil {
		return &model.CreateResponse{
			Error:       true,
			Description: fmt.Sprintf("Error to add header info: %q", err),
		}, err
	}

	trailer := metadata.Pairs("some-key", "some-value")
	err = grpc.SetTrailer(ctx, trailer)
	if err != nil {
		return &model.CreateResponse{
			Error:       true,
			Description: fmt.Sprintf("Error to add trailer info: %q", err),
		}, err
	}

	resp := &model.CreateResponse{
		Error:       false,
		Description: "Todo created.",
	}
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create a tcp listener: %q", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureValidToken),
	}

	serverRegistrar := grpc.NewServer(opts...)
	model.RegisterTaskerServer(serverRegistrar, &taskerServer{})

	log.Println("Listen:", listen.Addr())
	err = serverRegistrar.Serve(listen)
	if err != nil {
		log.Fatalf("Impossible to start up the server: %q", err)
	}
}

func ensureValidToken(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}

	return handler(ctx, req)
}

func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	return token == "some-secret-token"
}

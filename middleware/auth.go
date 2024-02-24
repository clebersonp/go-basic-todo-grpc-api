package middleware

import (
	"context"
	"github.com/clebersonp/go-basic-todo-grpc-api/failure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
)

func EnsureValidToken(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Println("Info:", info)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, failure.ErrMissingMetadata
	}

	if !valid(md["authorization"]) {
		return nil, failure.ErrInvalidToken
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

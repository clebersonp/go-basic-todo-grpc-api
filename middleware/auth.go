package middleware

import (
	"context"
	"github.com/clebersonp/go-basic-todo-grpc-api/failure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
)

const bearerTokenSecret = "some-secret-token"

func EnsureValidToken(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Println("Info:", info)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		e := failure.ErrMissingMetadata
		log.Println(e)
		return nil, e
	}

	if !valid(md["authorization"]) {
		e := failure.ErrInvalidToken
		log.Println(e)
		return nil, e
	}

	return handler(ctx, req)
}

func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	return token == bearerTokenSecret
}

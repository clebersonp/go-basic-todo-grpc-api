package middleware

import (
	"context"
	"github.com/clebersonp/go-basic-todo-grpc-api/failure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
)

const (
	bearerTokenSecret  = "some-secret-token"
	bypassSignUpMethod = "/Users/SignUp"
	bypassSignInMethod = "/Users/SignIn"
)

func EnsureValidToken(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	method := info.FullMethod
	log.Println("Token Validation Method:", method)
	if bypassSignUpMethod == method || bypassSignInMethod == method {
		log.Println("Bypass Method:", method)
		return handler(ctx, req)
	}

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

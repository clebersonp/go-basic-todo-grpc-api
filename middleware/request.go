package middleware

import (
	"context"
	"github.com/clebersonp/go-basic-todo-grpc-api/header"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RequestID(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// sending metadata to client
	reqId := uuid.New().String()
	md := metadata.Pairs(header.XRequestID, reqId)
	ctx = metadata.NewOutgoingContext(ctx, md)
	err = grpc.SendHeader(ctx, md)
	if err != nil {
		return nil, err
	}

	trailer := metadata.Pairs("some-key", "some-value")
	err = grpc.SetTrailer(ctx, trailer)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

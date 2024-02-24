package failure

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrMissingMetadata = status.Error(codes.InvalidArgument, "missing metadata")
	ErrInvalidToken    = status.Error(codes.Unauthenticated, "invalid token")
)

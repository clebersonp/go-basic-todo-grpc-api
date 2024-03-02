package failure

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrMissingMetadata        = status.Error(codes.InvalidArgument, "missing metadata")
	ErrInvalidToken           = status.Error(codes.Unauthenticated, "invalid token")
	ErrInvalidTodoID          = status.Error(codes.InvalidArgument, "invalid todo id")
	ErrTodoIdRequired         = status.Error(codes.InvalidArgument, "todo id is required")
	ErrTodoRequired           = status.Error(codes.InvalidArgument, "todo is required")
	ErrTodoNotFound           = status.Error(codes.NotFound, "todo not found")
	ErrEmailPasswordIncorrect = status.Error(codes.Unauthenticated, "your email or password is incorrect")
	ErrUserAlreadyExists      = status.Error(codes.AlreadyExists, "a user with that email address already exists")
)

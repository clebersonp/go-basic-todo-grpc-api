package user

import (
	"context"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto/user"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *ServerUser) SignUp(_ context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {

	// TODO must be implement
	log.Println("SignUp Unimplemented")

	return &emptypb.Empty{}, nil
}

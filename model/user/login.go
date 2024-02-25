package user

import (
	"context"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto/user"
	"log"
)

func (s *ServerUser) SignIn(_ context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	// TODO must implement
	log.Println("SignIn Unimplemented")

	return &pb.LoginUserResponse{
		Token:     "Unimplemented",
		ExpiresIn: 0,
	}, nil
}

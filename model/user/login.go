package user

import (
	"context"
	"github.com/clebersonp/go-basic-todo-grpc-api/header"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto/user"
	"log"
)

func (s *ServerUser) SignIn(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	reqId := header.GetRequestID(ctx)
	userEmail := req.GetEmail()
	log.Printf("%s: %s, User_Email: %s\n", header.XRequestID, reqId, userEmail)

	// TODO must implements token
	user, err := authentication(req.Email, req.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.LoginUserResponse{
		Token:     user.Id,
		ExpiresIn: 0,
	}, nil
}

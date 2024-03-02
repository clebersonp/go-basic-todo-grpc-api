package user

import (
	"context"
	"github.com/clebersonp/go-basic-todo-grpc-api/header"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto/user"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (s *ServerUser) SignUp(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	reqId := header.GetRequestID(ctx)
	log.Printf("User %v for X-Request-ID %q will be created", req, reqId)

	newUser := &pb.User{
		Id:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: timestamppb.Now(),
	}

	err := add(newUser)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	md := metadata.Pairs(header.UserId, newUser.Id)
	err = grpc.SetTrailer(ctx, md)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

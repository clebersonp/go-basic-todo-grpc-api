package model

import (
	"context"
	"github.com/clebersonp/go-basic-todo-grpc-api/header"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (t *TaskerServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.Response, error) {
	reqId := header.GetRequestID(ctx)
	log.Printf("Todo %v for X-Request-ID %q will be created", req, reqId)

	now := timestamppb.Now()
	newTodo := &pb.Todo{
		Id:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		Status:      pb.Status_CREATED,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	add(newTodo)

	md := metadata.Pairs(header.TodoID, newTodo.Id)
	err := grpc.SetTrailer(ctx, md)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp := &pb.Response{
		Error:       false,
		Description: "Todo created.",
	}
	return resp, nil
}

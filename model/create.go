package model

import (
	"context"
	"github.com/clebersonp/go-basic-todo-grpc-api/header"
	"github.com/clebersonp/go-basic-todo-grpc-api/proto"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (t *TaskerServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	reqId := header.GetRequestID(ctx)
	log.Printf("Todo %v for X-Request-ID %q will be created", req, reqId)

	now := timestamppb.Now()
	newTodo := &proto.Todo{
		Id:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		Status:      created,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	add(newTodo)

	resp := &proto.CreateResponse{
		Error:       false,
		Description: "Todo created.",
	}
	return resp, nil
}

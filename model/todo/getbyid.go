package model

import (
	"context"
	"github.com/clebersonp/go-basic-todo-grpc-api/failure"
	"github.com/clebersonp/go-basic-todo-grpc-api/header"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto"
	"github.com/clebersonp/go-basic-todo-grpc-api/validation"
	"log"
)

func (t *TaskerServer) GetByID(ctx context.Context, req *pb.TodoByIdRequest) (*pb.GetTodoByIdResponse, error) {
	reqId := header.GetRequestID(ctx)
	todoID := req.GetId()
	log.Printf("%s: %s, TodoID: %s\n", header.XRequestID, reqId, todoID)

	err := validation.TodoId(todoID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	todo, ok := getById(todoID)
	if !ok {
		e := failure.ErrTodoNotFound
		log.Println(e)
		return nil, e
	}
	return &pb.GetTodoByIdResponse{Todo: todo}, nil
}

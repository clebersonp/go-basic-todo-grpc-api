package model

import (
	"context"
	"github.com/clebersonp/go-basic-todo-grpc-api/failure"
	"github.com/clebersonp/go-basic-todo-grpc-api/header"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto"
	"github.com/google/uuid"
	"log"
)

func (t *TaskerServer) GetByID(ctx context.Context, req *pb.GetRequest) (*pb.Todo, error) {
	reqId := header.GetRequestID(ctx)
	todoID := req.GetId()
	log.Printf("%s: %s, TodoID: %s\n", header.XRequestID, reqId, todoID)

	err := validateId(todoID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	todo, ok := get(todoID)
	if !ok {
		e := failure.ErrTodoNotFound
		log.Println(e)
		return nil, e
	}
	return todo, nil
}

func validateId(id string) error {
	if id == "" {
		return failure.ErrTodoIdRequired
	}
	if _, err := uuid.Parse(id); err != nil {
		return failure.ErrInvalidTodoID
	}
	return nil
}

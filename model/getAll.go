package model

import (
	"context"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *TaskerServer) GetAll(ctx context.Context, empty *emptypb.Empty) (*pb.TodoListResponse, error) {
	var list []*pb.Todo
	for _, v := range todos {
		list = append(list, v)
	}
	return &pb.TodoListResponse{
		Count: int64(len(list)),
		Todos: list,
	}, nil
}

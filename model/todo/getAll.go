package todo

import (
	"context"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto/todo"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *TaskerServer) GetAll(_ context.Context, _ *emptypb.Empty) (*pb.TodoListResponse, error) {
	list := getAll()
	return &pb.TodoListResponse{
		Count: int64(len(list)),
		Todos: list,
	}, nil
}

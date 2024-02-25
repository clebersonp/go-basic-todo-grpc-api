package todo

import (
	"context"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto/todo"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (t *TaskerServer) Update(_ context.Context, req *pb.TodoUpdateRequest) (*emptypb.Empty, error) {
	todo := &pb.Todo{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}
	if err := update(todo); err != nil {
		log.Println(err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

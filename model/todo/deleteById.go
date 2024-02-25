package todo

import (
	"context"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto/todo"
	"github.com/clebersonp/go-basic-todo-grpc-api/validation"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (t *TaskerServer) DeleteByID(_ context.Context, req *pb.TodoByIdRequest) (*emptypb.Empty, error) {
	err := validation.TodoId(req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	deleteById(req.Id)
	return &emptypb.Empty{}, nil
}

package model

import "github.com/clebersonp/go-basic-todo-grpc-api/proto"

const (
	created  = "CREATED"
	canceled = "CANCELED"
	done     = "DONE"
)

var todos []*proto.Todo

type TaskerServer struct {
	proto.UnimplementedTaskerServer
}

func add(todo *proto.Todo) {
	todos = append(todos, todo)
}

package model

import "github.com/clebersonp/go-basic-todo-grpc-api/proto"

const (
	created  = "CREATED"
	canceled = "CANCELED"
	done     = "DONE"
)

var todos = make(map[string]*proto.Todo)

type TaskerServer struct {
	proto.UnimplementedTaskerServer
}

func add(todo *proto.Todo) {
	todos[todo.Id] = todo
}

func get(id string) (*proto.Todo, bool) {
	todo, ok := todos[id]
	return todo, ok
}

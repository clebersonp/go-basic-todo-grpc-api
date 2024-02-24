package model

import pb "github.com/clebersonp/go-basic-todo-grpc-api/proto"

var todos = make(map[string]*pb.Todo)

type TaskerServer struct {
	pb.UnimplementedTaskerServer
}

func add(todo *pb.Todo) {
	todos[todo.Id] = todo
}

func get(id string) (*pb.Todo, bool) {
	todo, ok := todos[id]
	return todo, ok
}

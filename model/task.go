package model

import (
	"github.com/clebersonp/go-basic-todo-grpc-api/failure"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto"
	"github.com/clebersonp/go-basic-todo-grpc-api/validation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var todos = make(map[string]*pb.Todo)

type TaskerServer struct {
	pb.UnimplementedTaskerServer
}

func add(todo *pb.Todo) {
	todos[todo.Id] = todo
}

func getById(id string) (*pb.Todo, bool) {
	todo, ok := todos[id]
	return todo, ok
}

func getAll() []*pb.Todo {
	var list []*pb.Todo
	for _, v := range todos {
		list = append(list, v)
	}
	return list
}

func deleteById(id string) {
	delete(todos, id)
}

func update(todo *pb.Todo) error {
	if todo == nil {
		return failure.ErrTodoRequired
	}
	if err := validation.TodoId(todo.Id); err != nil {
		return err
	}

	todoDb, ok := getById(todo.Id)
	if !ok {
		return failure.ErrTodoNotFound
	}

	todoDb.Title = todo.Title
	todoDb.Description = todo.Description
	todoDb.Status = todo.Status
	todoDb.UpdatedAt = timestamppb.Now()

	return nil
}

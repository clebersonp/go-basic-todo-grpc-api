package validation

import (
	"github.com/clebersonp/go-basic-todo-grpc-api/failure"
	"github.com/google/uuid"
)

func TodoId(id string) error {
	if id == "" {
		return failure.ErrTodoIdRequired
	}
	if _, err := uuid.Parse(id); err != nil {
		return failure.ErrInvalidTodoID
	}
	return nil
}

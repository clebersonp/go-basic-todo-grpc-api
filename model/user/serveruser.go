package user

import pb "github.com/clebersonp/go-basic-todo-grpc-api/proto/user"

type ServerUser struct {
	pb.UnimplementedUsersServer
}

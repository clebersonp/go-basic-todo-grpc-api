package user

import (
	"github.com/clebersonp/go-basic-todo-grpc-api/failure"
	pb "github.com/clebersonp/go-basic-todo-grpc-api/proto/user"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var users = make(map[string]*pb.User)

type ServerUser struct {
	pb.UnimplementedUsersServer
}

func add(user *pb.User) error {
	_, ok := getByEmail(user.Email)
	if ok {
		return failure.ErrUserAlreadyExists
	}
	hashedPassword, err := encryptPassword(user.Password)
	if err != nil {
		log.Println("Adding user Error:", err)
		return err
	}
	user.Password = hashedPassword
	users[user.Email] = user
	return nil
}

func getById(id string) (user *pb.User, ok bool) {
	for _, v := range users {
		if v.Id == id {
			user = v
			break
		}
	}
	return user, user != nil
}

func getAll() []*pb.User {
	var list []*pb.User
	for _, v := range users {
		list = append(list, v)
	}
	return list
}

func authentication(email, password string) (*pb.User, error) {
	user, ok := getByEmail(email)
	if !ok {
		return nil, failure.ErrEmailPasswordIncorrect
	}

	err := compareHashAndPassword(user.Password, password)
	if err != nil {
		log.Println("CompareHashAndPassword on login:", err.Error())
		return nil, failure.ErrEmailPasswordIncorrect
	}
	return user, nil
}

func getByEmail(email string) (user *pb.User, ok bool) {
	user, ok = users[email]
	return user, ok
}

func encryptPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pass), err
}

func compareHashAndPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

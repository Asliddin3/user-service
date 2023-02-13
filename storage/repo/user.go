package repo

import (
	"context"

	pb "github.com/Asliddin3/user-servis/genproto/user"
)

type UserStorageI interface {
	CreateUser(*pb.UserRequest) (*pb.UserResponse, error)
	GetAllUsers(*pb.Empty) (*pb.UsersResponse, error)
	GetUserInfo(context.Context, *pb.UserId) (*pb.UserResponse, error)
}

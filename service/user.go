package service

import (
	"context"

	pb "github.com/Asliddin3/user-servis/genproto/user"

	l "github.com/Asliddin3/user-servis/pkg/logger"
	"github.com/Asliddin3/user-servis/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService this struct have methods of user-service
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewUserService This func get user info
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

//GetUser This func get user info
func (s *UserService) GetUser(ctx context.Context, req *pb.UserId) (*pb.UserResponse, error) {
	user, err := s.storage.User().GetUserInfo(ctx, req)
	if err != nil {
		s.logger.Error("error while getting user", l.Any("error getting user", err))
		return nil, err
	}
	return user, nil
}

//GetAllUsers This func get all users info
func (s *UserService) GetAllUsers(ctx context.Context, req *pb.Empty) (*pb.UsersResponse, error) {
	users, err := s.storage.User().GetAllUsers(req)
	if err != nil {
		s.logger.Error("error while getting list user", l.Any("error getting user", err))
		return nil, err
	}
	return users, nil
}

//CreateUser This func creates new user
func (s *UserService) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := s.storage.User().CreateUser(req)
	if err != nil {
		s.logger.Error("error while creating user", l.Any("error creating user", err))
		return &pb.UserResponse{}, status.Error(codes.Internal, "something went wrong")
	}
	return user, nil
}

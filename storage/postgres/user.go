package postgres

import (
	"context"

	pb "github.com/Asliddin3/user-servis/genproto/user"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo create connection to postgres
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) GetUserInfo(ctx context.Context, req *pb.UserId) (*pb.UserResponse, error) {
	userResp := &pb.UserResponse{
		Id: req.Id,
	}
	err := r.db.QueryRow(`
	select name,age,phone
	from users
	where id=$1
	`, req.Id).Scan(
		&userResp.Name,
		&userResp.Age,
		&userResp.Phone)
	if err != nil {
		return nil, err
	}
	return userResp, nil
}

func (r *userRepo) GetAllUsers(req *pb.Empty) (*pb.UsersResponse, error) {
	arrUsers := &pb.UsersResponse{}
	rows, err := r.db.Query(`
	select  id,name,age,phone
	from users`)
	for rows.Next() {
		userInfo := &pb.UserResponse{}
		err = rows.Scan(
			&userInfo.Id,
			&userInfo.Name,
			&userInfo.Age,
			&userInfo.Phone)
		if err != nil {
			return nil, err
		}
		arrUsers.Users = append(arrUsers.Users, userInfo)
	}
	if err != nil {
		return nil, err
	}
	return arrUsers, nil
}

func (r *userRepo) CreateUser(req *pb.UserRequest) (*pb.UserResponse, error) {
	response := &pb.UserResponse{}
	err := r.db.QueryRow(`
	insert into users(name,age,phone)
	values($1,$2,$3)
	returning id,name,age,phone
	`, req.Name, req.Age, req.Phone).Scan(&response.Id,
		&response.Name,
		&response.Age,
		&response.Phone)
	if err != nil {
		return nil, err
	}
	return response, nil
}

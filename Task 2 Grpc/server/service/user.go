package service

import (
	"context"
	"fmt"
	"grpc-server/models"
	"grpc-server/pb"
	"grpc-server/repository"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func GetService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponce, error) {

	var user models.User

	user.Name = req.Name
	user.Age = int(req.Age)

	err := repository.CreateUser(ctx, user)

	if err != nil {
		return &pb.CreateUserResponce{
			Message: "Failed to create user",
		}, nil
	}

	return &pb.CreateUserResponce{Message: "User Created Succesfully"}, nil
}

func (u *UserService) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.Users, error) {
	users, err := repository.GetAllUsers(ctx)

	fmt.Println(users)

	userdata := make([]*pb.User, len(users))

	for i, v := range users {
		user := pb.User{
			Name: v.Name,
			Age:  uint32(v.Age),
		}

		userdata[i] = &user
	}

	if err != nil {
		return &pb.Users{
			Message: "Failed to get user data" + err.Error(),
		}, nil
	}

	fmt.Println(userdata)

	return &pb.Users{
		Message: "User data fetched succesfully",
		Users:   userdata,
	}, nil
}

func (u *UserService) GetUserByName(ctx context.Context, req *pb.GetUserByNameRequest) (*pb.User, error) {
	user, err := repository.GetUserByName(ctx, req.Name)

	if err != nil {
		return &pb.User{}, err
	}

	return &pb.User{
		Name: user.Name,
		Age:  uint32(user.Age),
	}, nil
}

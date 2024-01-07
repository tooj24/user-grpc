package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"grpc-clean/internal/models"
	interfaces "grpc-clean/pkg/v1"
	pb "grpc-clean/proto"
)

type UserService struct {
	useCase interfaces.UserUseCase
	pb.UnimplementedUserServiceServer
}

func NewServer(grpcServer *grpc.Server, useCase interfaces.UserUseCase) {
	userGrpc := &UserService{useCase: useCase}
	pb.RegisterUserServiceServer(grpcServer, userGrpc)
}

func (service *UserService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {
	data := service.transformUserRPC(req)
	if data.Email == "" || data.Name == "" {
		return &pb.UserProfileResponse{}, errors.New("please provide all fields")
	}

	user, err := service.useCase.Create(data)
	if err != nil {
		return &pb.UserProfileResponse{}, err
	}

	return service.transformUserModel(user), nil
}

func (service *UserService) Read(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	id := req.GetId()
	if id == "" {
		return &pb.UserProfileResponse{}, errors.New("id cannot be blank")
	}

	user, err := service.useCase.Get(id)
	if err != nil {
		return &pb.UserProfileResponse{}, err
	}

	return service.transformUserModel(user), nil
}

func (service *UserService) transformUserRPC(req *pb.CreateUserRequest) models.User {
	return models.User{
		Name:  req.Name,
		Email: req.Email,
	}
}

func (service *UserService) transformUserModel(user models.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{
		Id:    string(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
}

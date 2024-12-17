package internal

import (
	"context"
	"fmt"

	userv1 "go-server/gen/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	userv1.UnimplementedUserServiceServer
	users map[string]*userv1.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*userv1.User),
	}
}

func (s *UserService) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	user, exists := s.users[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "user not found: %s", req.Id)
	}

	return &userv1.GetUserResponse{User: user}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	// 간단한 ID 생성 (실제 프로젝트에서는 UUID 등을 사용해야 합니다)
	id := fmt.Sprintf("user_%d", len(s.users)+1)

	user := &userv1.User{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	s.users[id] = user

	return &userv1.CreateUserResponse{User: user}, nil
}

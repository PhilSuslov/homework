package service

import (
	"context"

	"github.com/PhilSuslov/homework/iam/internal/model"
)


type AuthService interface {
	Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error) 
	Whoami(ctx context.Context, req model.WhoamiRequest) (*model.WhoamiResponse, error)
}

type UserService interface {
	GetUser(ctx context.Context, req model.GetUserRequest) (model.GetUserResponse, error)
	Register(ctx context.Context, req model.RegisterRequest) (model.RegisterResponse, error)
}
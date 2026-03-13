package user

import (
	"context"

	conv "github.com/PhilSuslov/homework/iam/internal/converter"
	"github.com/PhilSuslov/homework/iam/internal/model"
)

func (s *Service) GetUser(ctx context.Context, req model.GetUserRequest) (model.GetUserResponse, error){
	user, err := s.repo.Get(ctx, req.UserUuid)
	if err != nil{
		return model.GetUserResponse{}, model.ErrFailGetUser
	}

	return conv.UserToGetUserResponse(user), nil
}
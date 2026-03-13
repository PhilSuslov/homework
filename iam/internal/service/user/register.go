package user

import (
	"context"

	conv "github.com/PhilSuslov/homework/iam/internal/converter"
	"github.com/PhilSuslov/homework/iam/internal/model"
)

func (s *Service) Register(ctx context.Context, req model.RegisterRequest) (model.RegisterResponse, error) {
	user, err := s.repo.Create(ctx, conv.UserToRepo(req.User))
	if err != nil {
		return model.RegisterResponse{}, model.ErrFailRegister
	}

	return conv.StringToRegisterResponse(user), nil
}

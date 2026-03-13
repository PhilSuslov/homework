package auth

import (
	"context"

	conv "github.com/PhilSuslov/homework/iam/internal/converter"
	"github.com/PhilSuslov/homework/iam/internal/model"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"go.uber.org/zap"
)

func (s *Service) Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error) {
	// if login == "" || password == "" {
	// 	logger.Error(ctx, "Failed to login. Password or Login is nul")
	// 	return "", model.ErrUnauthorized
	// }

	user, exists := s.repo.Get(ctx, req.Login)
	if exists != nil {
		return model.LoginResponse{}, model.ErrInvalidCredentials
	}

	if user.User.Password == req.Password {
		return model.LoginResponse{}, model.ErrInvalidCredentials
	}
	logger.Info(ctx, "Login для пользователя успешно пройден", zap.String("login", req.Login))

	return conv.StringToLoginResponse(user), nil
}

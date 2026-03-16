package auth

import (
	"context"
	"fmt"
	"time"

	conv "github.com/PhilSuslov/homework/iam/internal/converter"
	"github.com/PhilSuslov/homework/iam/internal/model"
	repoModel "github.com/PhilSuslov/homework/iam/internal/repository/model"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *Service) Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error) {
	userUUID, exists := s.pgRepo.Login(ctx, req.Login, req.Password)
	if exists != nil {
		return model.LoginResponse{}, model.ErrInvalidCredentials
	}

	user, err := s.pgRepo.Get(ctx, userUUID)
	if err != nil {
		return model.LoginResponse{}, model.ErrInvalidCredentials
	}

	fmt.Println("Service -> auth -> login -> user: ", user)

	sessionUUID := uuid.NewString()
	session := repoModel.Session{
		Uuid: sessionUUID,
		User: user,
	}

	if err := s.redisRepo.Set(ctx, sessionUUID, session, 24*time.Hour); err != nil {
		return model.LoginResponse{}, err
	}

	if err := s.redisRepo.AddSessionToUserSet(ctx, userUUID, sessionUUID, 24*time.Hour); err != nil {
		return model.LoginResponse{}, err
	}

	logger.Info(ctx, "Login для пользователя успешно пройден", zap.String("login", req.Login))

	return conv.StringToLoginResponse(sessionUUID), nil
}

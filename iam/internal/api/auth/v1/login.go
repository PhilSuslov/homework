package v1

import (
	"context"

	conv "github.com/PhilSuslov/homework/iam/internal/converter"
	"github.com/PhilSuslov/homework/iam/internal/model"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	authV1 "github.com/PhilSuslov/homework/shared/pkg/proto/auth/v1"
)

func (a *api) Login(ctx context.Context, req *authV1.LoginRequest) (*authV1.LoginResponser, error) {
	if req.Login == "" || req.Password == "" {
		logger.Error(ctx, "Failed to login. Password or Login is nul")
		return nil, model.ErrUnauthorized
	}

	login, err := a.iam.Login(ctx, conv.ProtoToLoginRequestModel(req))
	if err != nil {
		return nil, model.ErrInvalidCredentials
	}

	return &authV1.LoginResponser{
		SessionUuid: login.SessionUuid,
	}, nil
}

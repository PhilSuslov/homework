package v1

import (
	"context"

	conv "github.com/PhilSuslov/homework/iam/internal/converter"
	"github.com/PhilSuslov/homework/iam/internal/model"
	userV1 "github.com/PhilSuslov/homework/shared/pkg/proto/common/v1"
)

func (a *api) Register(ctx context.Context, req *userV1.RegisterRequest) (*userV1.RegisterResponser, error) {
	if req.User.Login == "" || req.User.Password == "" {
		return nil, model.ErrFailRegister
	}

	res, err := a.user.Register(ctx, conv.ProtoToRegister(req))
	if err != nil {
		return nil, model.ErrFailRegister
	}

	return &userV1.RegisterResponser{
		Uuid: res.Uuid,
	}, nil
}

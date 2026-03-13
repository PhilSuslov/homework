package v1

import (
	"context"

	conv "github.com/PhilSuslov/homework/iam/internal/converter"
	"github.com/PhilSuslov/homework/iam/internal/model"
	authV1 "github.com/PhilSuslov/homework/shared/pkg/proto/auth/v1"
)

func (a *api) Whoami(ctx context.Context, req *authV1.WhoamiRequest) (*authV1.WhoamiResponser, error) {
	if req.SessionUuid == "" {
		return nil, model.ErrSessionNotFound
	}

	res, err := a.iam.Whoami(ctx, conv.ProtoToWhoamiRequestModel(req))
	if err != nil {
		return nil, model.ErrSessionNotFound
	}

	return &authV1.WhoamiResponser{
		UserUuid: res.UserUuid,
		Login:    res.Login,
		Email:    res.Email,
	}, nil
}

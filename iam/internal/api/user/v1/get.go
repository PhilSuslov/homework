package v1

import (
	"context"

	conv "github.com/PhilSuslov/homework/iam/internal/converter"
	"github.com/PhilSuslov/homework/iam/internal/model"
	userV1 "github.com/PhilSuslov/homework/shared/pkg/proto/common/v1"
)

func (a *api) GetUser(ctx context.Context, req *userV1.GetUserRequest) (*userV1.GetUserResponser, error) {
	if req.UserUuid == "" {
		return nil, model.ErrFailGetUser
	}

	user, err := a.user.GetUser(ctx, conv.ProtoToGetUserRequest(req))
	if err != nil {
		return nil, model.ErrFailGetUser
	}

	return &userV1.GetUserResponser{
		User: conv.GetUserResponseToProto(user.User),
	}, nil
}

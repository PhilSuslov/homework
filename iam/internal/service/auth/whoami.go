package auth

import (
	"context"

	conv "github.com/PhilSuslov/homework/iam/internal/converter"
	"github.com/PhilSuslov/homework/iam/internal/model"
)

func (s *Service) Whoami(ctx context.Context, req model.WhoamiRequest) (*model.WhoamiResponse, error) {
	session, err := s.repo.Get(ctx, req.SessionUuid)
	if err != nil {
		return nil, model.ErrSessionNotFound
	}

	user, err := s.repo.Get(ctx, session.User.User_uuid)
	if err != nil {
		return nil, err
	}

	return conv.UserAuthToWhoamiResponse(user), err
}

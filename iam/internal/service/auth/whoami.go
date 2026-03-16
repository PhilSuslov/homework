package auth

import (
	"context"
	"fmt"

	conv "github.com/PhilSuslov/homework/iam/internal/converter"
	"github.com/PhilSuslov/homework/iam/internal/model"
)

func (s *Service) Whoami(ctx context.Context, req model.WhoamiRequest) (*model.WhoamiResponse, error) {
	fmt.Println("Service -> auth -> whoami: ",req.SessionUuid)
	session, err := s.redisRepo.Get(ctx, req.SessionUuid)
	fmt.Println(session, err)
	if err != nil {
		return nil, model.ErrSessionNotFound
	}

	fmt.Println("Service -> auth -> whoami -> session.conv: ", conv.UserAuthToWhoamiResponse(session))

	return conv.UserAuthToWhoamiResponse(session), err
}

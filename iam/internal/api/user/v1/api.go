package v1

import (
	"github.com/PhilSuslov/homework/iam/internal/service"	
	userV1 "github.com/PhilSuslov/homework/shared/pkg/proto/common/v1"
)

type api struct {
	userV1.UnimplementedUserServiceServer

	user service.UserService
}

func NewAPI(user service.UserService) *api {
	return &api{
		user: user,
	}
}
package v1

import (
	"github.com/PhilSuslov/homework/iam/internal/service"	
	authV1 "github.com/PhilSuslov/homework/shared/pkg/proto/auth/v1"
)

type api struct {
	authV1.UnimplementedAuthServiceServer

	iam service.AuthService
}

func NewAPI(iam service.AuthService) *api {
	return &api{
		iam: iam,
	}
}
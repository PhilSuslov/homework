package user

import (
	repo "github.com/PhilSuslov/homework/iam/internal/repository"
	def "github.com/PhilSuslov/homework/iam/internal/service"

)

var _ def.UserService = (*Service)(nil)

type Service struct {
	repo repo.IAMPostgresRepository
}

func NewService(repo repo.IAMPostgresRepository) *Service {
	return &Service{repo: repo}
}

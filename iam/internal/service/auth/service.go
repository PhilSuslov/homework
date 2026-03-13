package auth

import (

	repo "github.com/PhilSuslov/homework/iam/internal/repository"
	def "github.com/PhilSuslov/homework/iam/internal/service"
)
var _ def.AuthService = (*Service)(nil)

type Service struct {
	repo repo.IAMRedisRepository
}

func NewService(repo repo.IAMRedisRepository) *Service {
	return &Service{repo: repo}
}

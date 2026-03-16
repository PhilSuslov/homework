package auth

import (
	repo "github.com/PhilSuslov/homework/iam/internal/repository"
	def "github.com/PhilSuslov/homework/iam/internal/service"
)

var _ def.AuthService = (*Service)(nil)

type Service struct {
	redisRepo repo.IAMRedisRepository
	pgRepo    repo.IAMPostgresRepository
}

func NewService(redisRepo repo.IAMRedisRepository, pgRepo repo.IAMPostgresRepository) *Service {
	return &Service{redisRepo: redisRepo,
		pgRepo: pgRepo}
}

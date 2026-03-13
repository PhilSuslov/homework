package session

import (
	"github.com/PhilSuslov/homework/platform/pkg/cache"
	def "github.com/PhilSuslov/homework/iam/internal/repository"

)

var _ def.IAMRedisRepository = (*Repository)(nil)

type Repository struct {
	cache cache.RedisClient
}

func NewRepository(cache cache.RedisClient) *Repository {
	return &Repository{cache: cache}
}

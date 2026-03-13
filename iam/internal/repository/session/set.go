package session

import (
	"context"
	"time"

	repoConverter "github.com/PhilSuslov/homework/iam/internal/repository/converter"

	"github.com/PhilSuslov/homework/iam/internal/repository/model"
)

func (r *Repository) Set(ctx context.Context, uuid string, session model.Session, ttl time.Duration) error {
	cacheKey := r.getCacheKey(uuid)

	redisView := repoConverter.SessionFromRedis(session)

	err := r.cache.HashSet(ctx, cacheKey, redisView)
	if err != nil {
		return err
	}

	return r.cache.Expire(ctx, cacheKey, ttl)
}

package session

import (
	"context"
	"errors"
	"fmt"

	repoConverter "github.com/PhilSuslov/homework/iam/internal/repository/converter"
	repoModel "github.com/PhilSuslov/homework/iam/internal/repository/model"
	redigo "github.com/gomodule/redigo/redis"

	"github.com/PhilSuslov/homework/iam/internal/model"
)

const (
	cachePrefix = "iam:session:"
)

func (r Repository) getCacheKey(uuid string) string {
	return fmt.Sprintf("%s%s", cachePrefix, uuid)
}

func (r *Repository) Get(ctx context.Context, uuid string) (model.Session, error) {
	cacheKey := r.getCacheKey(uuid)
	fmt.Println("Get -> session ",cacheKey)

	values, err := r.cache.HGetAll(ctx, cacheKey)
	if err != nil {
		if errors.Is(err, redigo.ErrNil) {
			return model.Session{}, model.ErrSessionNotFound
		}
		return model.Session{}, err
	}
	fmt.Println("Session -> Get -> Values...:")
	fmt.Println(values...)

	if len(values) == 0 {
		return model.Session{}, model.ErrSessionNotFound
	}

	var sessionRedisView repoModel.Session
	err = redigo.ScanStruct(values, &sessionRedisView)
	if err != nil {
		return model.Session{}, err
	}

	fmt.Println("Session -> Get -> sessionRedisView: ",sessionRedisView)
	return repoConverter.SessionFromRedis(sessionRedisView), nil
}

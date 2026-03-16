package repository

import (
	"context"
	"time"

	repoModel"github.com/PhilSuslov/homework/iam/internal/repository/model"
	"github.com/PhilSuslov/homework/iam/internal/model"


)

type IAMPostgresRepository interface {
	Create(ctx context.Context, user repoModel.UserRedis) (string, error)
	Get(ctx context.Context, userUuid string) (repoModel.UserRedis, error) // Предполагаю, что тут должен быть User
	Login(ctx context.Context, login, password string) (string, error) 
}

type IAMRedisRepository interface {
	AddSessionToUserSet(ctx context.Context, userUUID, sessionUUID string, ttl time.Duration) error
	Set(ctx context.Context, uuid string, session repoModel.Session, ttl time.Duration) error
	Get(ctx context.Context, uuid string) (model.Session, error) // Предполагаю, что тут должен быть User
	Delete(ctx context.Context, uuid string) error // Может другой функционал(нет delete в session)
}
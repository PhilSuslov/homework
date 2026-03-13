package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/stdlib"

	"github.com/jackc/pgx/v5/pgxpool"

	apiAuthV1 "github.com/PhilSuslov/homework/iam/internal/api/auth/v1"
	apiUserV1 "github.com/PhilSuslov/homework/iam/internal/api/user/v1"
	"github.com/PhilSuslov/homework/iam/internal/config"
	repo "github.com/PhilSuslov/homework/iam/internal/repository"
	repoSession "github.com/PhilSuslov/homework/iam/internal/repository/session"
	repoUser "github.com/PhilSuslov/homework/iam/internal/repository/user"
	"github.com/PhilSuslov/homework/iam/internal/service"
	serviceAuth "github.com/PhilSuslov/homework/iam/internal/service/auth"
	serviceUser "github.com/PhilSuslov/homework/iam/internal/service/user"
	"github.com/PhilSuslov/homework/platform/pkg/cache"
	"github.com/PhilSuslov/homework/platform/pkg/cache/redis"
	"github.com/PhilSuslov/homework/platform/pkg/logger"
	authV1 "github.com/PhilSuslov/homework/shared/pkg/proto/auth/v1"
	userV1 "github.com/PhilSuslov/homework/shared/pkg/proto/common/v1"
	redigo "github.com/gomodule/redigo/redis"
)

type diContainer struct {
	authV1API authV1.AuthServiceServer
	userV1API userV1.UserServiceServer

	userService service.UserService
	authService service.AuthService

	userRepo repo.IAMPostgresRepository
	authRepo repo.IAMRedisRepository

	postgresConn *pgxpool.Pool

	redisPool   *redigo.Pool
	redisClient cache.RedisClient
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) AuthV1API(ctx context.Context) authV1.AuthServiceServer {
	if d.authV1API == nil {
		d.authV1API = apiAuthV1.NewAPI(d.PartAuthService(ctx))
	}
	return d.authV1API
}

func (d *diContainer) UserAuthV1API(ctx context.Context) userV1.UserServiceServer {
	if d.userV1API == nil {
		d.userV1API = apiUserV1.NewAPI(d.PartUserService(ctx))
	}
	return d.userV1API
}

func (d *diContainer) PartAuthService(ctx context.Context) service.AuthService {
	if d.authService == nil {
		d.authService = serviceAuth.NewService(d.PartAuthRepo(ctx))
	}
	return d.authService
}

func (d *diContainer) PartUserService(ctx context.Context) service.UserService {
	if d.userService == nil {
		d.userService = serviceUser.NewService(d.PartUserRepo(ctx))
	}
	return d.userService
}

func (d *diContainer) PartAuthRepo(ctx context.Context) repo.IAMRedisRepository {
	if d.authRepo == nil {
		d.authRepo = repoSession.NewRepository(d.RedisClient())

	}
	return d.authRepo
}

func (d *diContainer) PartUserRepo(ctx context.Context) repo.IAMPostgresRepository {
	if d.userRepo == nil {
		d.userRepo = repoUser.NewIAMRepo(d.PostgresPool(ctx))
	}
	return d.userRepo
}

func (d *diContainer) RedisClient() cache.RedisClient {
	if d.redisClient == nil {
		d.redisClient = redis.NewClient(d.RedisPool(), logger.Logger(), config.AppConfig().Redis.ConnectionTimeout())
	}

	return d.redisClient
}

func (d *diContainer) RedisPool() *redigo.Pool {
	if d.redisPool == nil {
		d.redisPool = &redigo.Pool{
			MaxIdle:     config.AppConfig().Redis.MaxIdle(),
			IdleTimeout: config.AppConfig().Redis.IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", config.AppConfig().Redis.Address())
			},
		}
	}
	return d.redisPool
}

func (d *diContainer) PostgresPool(ctx context.Context) *pgxpool.Pool {
	if d.postgresConn != nil {
		return d.postgresConn
	}

	conn, err := pgxpool.New(ctx, "postgres://demo:demo@localhost:5435/orders?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Проверяем доступность базы
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}

	// Миграции через Goose
	migrationsDir := "../migrations"
	db := stdlib.OpenDB(*conn.Config().ConnConfig)
	migrationsRunner := repoUser.NewMigrator(db, migrationsDir)

	err = migrationsRunner.Up()
	if err != nil {
		log.Fatalf("Ошибка миграции базы данных: %v", err)
	}

	d.postgresConn = conn
	return d.postgresConn
}

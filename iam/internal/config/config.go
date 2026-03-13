package config

import (
	"os"

	"github.com/PhilSuslov/homework/iam/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	IAMGRPC  IAMGRPCConfig
	Logger   LoggerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Session  SessionConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	iamgrpcCfg, err := env.NewIAMGRPCConfig()
	if err != nil {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	postgresCfg, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	redisCfg, err := env.NewRedisConfig()
	if err != nil {
		return err
	}

	sessionCfg, err := env.NewSessionConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		IAMGRPC:  iamgrpcCfg,
		Logger:   loggerCfg,
		Postgres: postgresCfg,
		Redis:    redisCfg,
		Session:  sessionCfg,
	}

	return nil

}

func AppConfig() *config {
	return appConfig
}

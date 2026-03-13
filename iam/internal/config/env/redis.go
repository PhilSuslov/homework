package env

import (
	"net"
	"time"

	"github.com/caarlos0/env/v11"
)

/*
IAM_REDIS_HOST=localhost
IAM_REDIS_PORT=6333
IAM_EXTERNAL_REDIS_PORT=6333
IAM_REDIS_CONNECTION_TIMEOUT=10s
IAM_REDIS_MAX_IDLE=10
IAM_REDIS_IDLE_TIMEOUT=10s
*/

type redisEnvConfig struct {
	Host string `env:"IAM_REDIS_HOST,required"`
	Port string `env:"IAM_REDIS_PORT,required"`
	ConnectionTimeout time.Duration `env:"IAM_REDIS_CONNECTION_TIMEOUT,required"`
	MaxIdle int `env:"IAM_REDIS_MAX_IDLE,required"`
	IdleTimeout time.Duration `env:"IAM_REDIS_IDLE_TIMEOUT,required"`
}

type redisConfig struct {
	raw redisEnvConfig
}

func NewRedisConfig() (*redisConfig, error) {
	var raw redisEnvConfig
	err := env.Parse(&raw)
	if err != nil{
		return nil, err
	}

	return &redisConfig{raw: raw}, nil
}

func (cfg *redisConfig) Address() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}

func (cfg *redisConfig) ConnectionTimeout() time.Duration {
	return cfg.raw.ConnectionTimeout
}

func (cfg *redisConfig) MaxIdle() int {
	return cfg.raw.MaxIdle
}

func (cfg *redisConfig) IdleTimeout() time.Duration {
	return cfg.raw.IdleTimeout
}

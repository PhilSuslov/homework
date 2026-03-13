package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type iamGRPCEnvConfig struct {
	Host string `env:"IAM_GRPC_HOST,required"`
	Port string `env:"IAM_GRPC_PORT,required"`
}

type IAMGRPCConfig struct {
	raw iamGRPCEnvConfig
}

func NewIAMGRPCConfig() (*IAMGRPCConfig, error) {
	var raw iamGRPCEnvConfig

	err := env.Parse(&raw)
	if err != nil {
		return nil, err
	}

	return &IAMGRPCConfig{raw: raw}, nil
}

func (cfg IAMGRPCConfig) Address() string{
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}
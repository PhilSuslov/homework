package env

import (
	"log"
	"net"

	"github.com/caarlos0/env/v11"
)

type orderHTTPEnvConfig struct {
	Host    string `env:"HTTP_HOST,required"`
	Port    string `env:"HTTP_PORT,required"`
	Timeout string `env:"HTTP_READ_TIMEOUT,required"`
}

type orderHTTPConfig struct {
	raw orderHTTPEnvConfig
}

func NewOrderHTTPConfig()(*orderHTTPConfig,error){
	var raw orderHTTPEnvConfig
	err := env.Parse(&raw)
	if err != nil{
		log.Printf("failed to parse OrderHTTP: %v\n", err)
		return nil, err
	}
	return &orderHTTPConfig{raw: raw}, nil
}

func (cfg *orderHTTPConfig) Address() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}

func (cfg *orderHTTPConfig) Timeout() string{
	return cfg.raw.Timeout
}

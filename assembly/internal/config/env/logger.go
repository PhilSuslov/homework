package env

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type loggerEnvConfig struct {
	Level  string `env:"LOGGER_LEVEL,required"`
	AsJSON bool   `env:"LOGGER_AS_JSON,required"`
}

type loggerConfig struct {
	raw loggerEnvConfig
}

func NewLoggerConfig() (*loggerConfig, error) {
	var raw loggerEnvConfig

	if err := env.Parse(&raw); err != nil {
		log.Fatal("Не инициализирован логгер через NewLoggerConfig")
		return nil, err
	}
	return &loggerConfig{raw: raw}, nil
}

func (cfg *loggerConfig) Level() string {
	return cfg.raw.Level
}

func (cfg *loggerConfig) AsJson() bool {
	return cfg.raw.AsJSON
}

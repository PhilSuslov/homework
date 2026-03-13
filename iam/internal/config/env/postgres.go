package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type postgresEnvConfig struct {
	Host string `env:"IAM_POSTGRES_HOST,required"`
	Port string `env:"IAM_POSTGRES_PORT,required"`
	// ExternalPort string `env:"IAM_EXTERNAL_POSTGRES_PORT,required"`
	User     string `env:"IAM_POSTGRES_USER,required"`
	Password string `env:"IAM_POSTGRES_PASSWORD,required"`
	Database string `env:"IAM_POSTGRES_DB,required"`
	SSLmode  string `env:"IAM_POSTGRES_SSL_MODE,required"`
	// Migration    string `env:"IAM_MIGRATION_DIRECTORY,required"`
}

/*
IAM_POSTGRES_HOST=localhost
IAM_POSTGRES_PORT=5444
IAM_EXTERNAL_POSTGRES_PORT=5444
IAM_POSTGRES_USER=iam_user
IAM_POSTGRES_PASSWORD=iam_password
IAM_POSTGRES_DB=iam
IAM_POSTGRES_SSL_MODE=disable
IAM_MIGRATION_DIRECTORY=./iam/migrations
*/

type postgresConfig struct {
	raw postgresEnvConfig
}

func NewPostgresConfig() (*postgresConfig, error) {
	var raw postgresEnvConfig
	err := env.Parse(&raw)
	if err != nil {
		return nil, err
	}
	return &postgresConfig{raw: raw}, nil
}

func (cfg *postgresConfig) URI() string {
	fmt.Printf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.raw.User,
		cfg.raw.Password,
		cfg.raw.Host,
		cfg.raw.Port,
		cfg.raw.Database,
		cfg.raw.SSLmode)

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.raw.User,
		cfg.raw.Password,
		cfg.raw.Host,
		cfg.raw.Port,
		cfg.raw.Database,
		cfg.raw.SSLmode,
	)
}

func (cfg *postgresConfig) DatabaseName() string {
	return cfg.raw.Database
}

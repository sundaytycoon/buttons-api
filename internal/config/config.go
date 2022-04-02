package config

import (
	"time"

	"github.com/pkg/errors"
)

type Config struct {
	Env     string `env:"ENV"`
	Debug   bool   `env:"DEBUG,default=false"`
	Release string // inject when in dockerized by git hash

	ApplicationHTTP struct {
		InternalDSN string        `env:"APP_HTTP_INTERNAL_DSN,default=localhost:40080"`
		ExternalDSN string        `env:"APP_HTTP_EXTERNAL_DSN"`
		Timeout     time.Duration `env:"APP_HTTP_TIMEOUT,default=2s"`
	}
	ApplicationGRPC struct {
		InternalDSN string        `env:"APP_GRPC_INTERNAL_DSN,default=localhost:40081"`
		ExternalDSN string        `env:"APP_GRPC_EXTERNAL_DSN"`
		Timeout     time.Duration `env:"APP_GRPC_TIMEOUT,default=2s"`
	}
	ApplicationSwagger struct {
		InternalDSN string        `env:"APP_SWAGGER_INTERNAL_DSN,default=localhost:40082"`
		ExternalDSN string        `env:"APP_SWAGGER_EXTERNAL_DSN"`
		Timeout     time.Duration `env:"APP_SWAGGER_TIMEOUT,default=2s"`
	}

	AdminServer struct {
		DSN     string        `env:"ADMIN_SERVER_DSN,default=localhost:43000"`
		Timeout time.Duration `env:"ADMIN_SERVER_TIMEOUT,default=2s"`
	}
	AdminWeb struct {
		DSN string `env:"ADMIN_WEB_DSN,default=localhost:43080"`
	}

	ServiceDB struct {
		DSN          string `env:"SERVICE_DB_DSN,required=true"`
		MaxIdleConns int    `env:"SERVICE_DB_MAX_IDLE_CONNS,required=true"`
		MaxOpenConns int    `env:"SERVICE_DB_MAX_OPEN_CONNS,required=true"`
	}

	Google struct {
		OAuthCallbackURL string `env:"GOOGLE_OAUTH_CALLBACK_URL,required=true"`
		ClientID         string `env:"GOOGLE_CLIENT_ID,required=true"`
		ClientSecret     string `env:"GOOGLE_CLIENT_SECRET,required=true"`
	}
}

func (c *Config) Self() (*Config, error) {
	if c == nil {
		return nil, errors.New("configuration struct is null")
	}
	return c, nil
}

package config

import (
	"os"

	"github.com/satanaroom/chat_server/internal/errs"
)

var _ AuthClientConfig = (*authClientConfig)(nil)

const authPortEnvName = "AUTH_PORT"

type AuthClientConfig interface {
	Port() string
}

type authClientConfig struct {
	port string
}

func NewAuthClientConfig() (*authClientConfig, error) {
	port := os.Getenv(authPortEnvName)
	if port == "" {
		return nil, errs.ErrAuthClientPortNotFound
	}

	return &authClientConfig{
		port: port,
	}, nil
}

func (c *authClientConfig) Port() string {
	return c.port
}

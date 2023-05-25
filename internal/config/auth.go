package config

import (
	"github.com/satanaroom/auth/pkg/env"
)

var _ AuthClientConfig = (*authClientConfig)(nil)

const authHostEnvName = "AUTH_HOST"

type AuthClientConfig interface {
	Host() string
}

type authClientConfig struct {
	host string
}

func NewAuthClientConfig() (*authClientConfig, error) {
	var host string
	env.ToString(&host, authHostEnvName, "localhost:50051")

	return &authClientConfig{
		host: host,
	}, nil
}

func (c *authClientConfig) Host() string {
	return c.host
}

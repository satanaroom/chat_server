package config

import (
	"github.com/satanaroom/auth/pkg/env"
)

var _ SwaggerConfig = (*swaggerConfig)(nil)

const swaggerHostEnvName = "SWAGGER_HOST"

type SwaggerConfig interface {
	Host() string
}

type swaggerConfig struct {
	host string
}

func NewSwaggerConfig() (*swaggerConfig, error) {
	var host string
	env.ToString(&host, swaggerHostEnvName, "localhost:8091")

	return &swaggerConfig{
		host: host,
	}, nil
}

func (c *swaggerConfig) Host() string {
	return c.host
}

package config

import (
	"github.com/satanaroom/auth/pkg/env"
)

var _ HTTPConfig = (*httpConfig)(nil)

const httpHostEnvName = "HTTP_HOST"

type HTTPConfig interface {
	Host() string
}

type httpConfig struct {
	host string
}

func NewHTTPConfig() (*httpConfig, error) {
	var host string
	env.ToString(&host, httpHostEnvName, "localhost:8081")

	return &httpConfig{
		host: host,
	}, nil
}

func (c *httpConfig) Host() string {
	return c.host
}

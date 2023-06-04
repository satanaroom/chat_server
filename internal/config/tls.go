package config

import (
	"github.com/satanaroom/auth/pkg/env"
)

var _ TLSConfig = (*tlsConfig)(nil)

const tlsCertFileEnvName = "TLS_AUTH_CERT_FILE"

type TLSConfig interface {
	CertFile() string
}

type tlsConfig struct {
	certFile string
}

func NewTLSConfig() (*tlsConfig, error) {
	var certFile string

	env.ToString(&certFile, tlsCertFileEnvName, "service.pem")

	return &tlsConfig{
		certFile: certFile,
	}, nil
}

func (c *tlsConfig) CertFile() string {
	return c.certFile
}

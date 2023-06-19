package config

import (
	"github.com/satanaroom/auth/pkg/env"
)

var _ TLSConfig = (*tlsConfig)(nil)

const (
	tlsAuthCertFileEnvName   = "TLS_AUTH_CERT_FILE"
	tlsServerCertFileEnvName = "TLS_SERVER_CERT_FILE"
	tlsServerKeyFileEnvName  = "TLS_SERVER_KEY_FILE"
)

type TLSConfig interface {
	AuthCertFile() string
	ServerCertFile() string
	ServerKeyFile() string
}

type tlsConfig struct {
	authCertFile   string
	serverCertFile string
	serverKeyFile  string
}

func NewTLSConfig() (*tlsConfig, error) {
	var authCertFile, serverCertFile, serverKeyFile string

	env.ToString(&authCertFile, tlsAuthCertFileEnvName, "service.pem")
	env.ToString(&serverCertFile, tlsServerCertFileEnvName, "service.pem")
	env.ToString(&serverKeyFile, tlsServerKeyFileEnvName, "service.key")

	return &tlsConfig{
		authCertFile:   authCertFile,
		serverCertFile: serverCertFile,
		serverKeyFile:  serverKeyFile,
	}, nil
}

func (c *tlsConfig) AuthCertFile() string {
	return c.authCertFile
}

func (c *tlsConfig) ServerCertFile() string {
	return c.serverCertFile
}

func (c *tlsConfig) ServerKeyFile() string {
	return c.serverKeyFile
}

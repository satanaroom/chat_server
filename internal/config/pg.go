package config

import (
	"os"

	"github.com/satanaroom/chat_server/internal/errs"
)

var _ PGConfig = (*pgConfig)(nil)

const pgDSNEnvName = "PG_DSN"

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func NewPGConfig() (*pgConfig, error) {
	dsn := os.Getenv(pgDSNEnvName)
	if len(dsn) == 0 {
		return nil, errs.ErrDSNNotFound
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}

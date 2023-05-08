package pg

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/satanaroom/auth/pkg/logger"
)

var _ Client = (*client)(nil)

type Client interface {
	Close() error
	PG() PG
}

type client struct {
	pg PG
}

func NewClient(ctx context.Context, pgCfg *pgxpool.Config) (*client, error) {
	dbc, err := pgxpool.ConnectConfig(ctx, pgCfg)
	if err != nil {
		logger.Fatalf("failed to get db connection: %s", err.Error())
	}

	return &client{
		pg: &pg{
			pool: dbc,
		},
	}, nil
}

func (c *client) PG() PG {
	return c.pg
}

func (c *client) Close() error {
	if c.pg != nil {
		return c.pg.Close()
	}
	return nil
}

package chat

import (
	"github.com/satanaroom/chat_server/internal/client/pg"
)

var _ Repository = (*repository)(nil)

type Repository interface {
}

type repository struct {
	pgClient pg.Client
}

func NewRepository(pgClient pg.Client) *repository {
	return &repository{
		pgClient: pgClient,
	}
}

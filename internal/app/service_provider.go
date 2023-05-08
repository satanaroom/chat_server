package app

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_server/internal/client/pg"
	"github.com/satanaroom/chat_server/internal/closer"
	"github.com/satanaroom/chat_server/internal/config"
	chatRepository "github.com/satanaroom/chat_server/internal/repository/chat"
	chatService "github.com/satanaroom/chat_server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig config.PGConfig

	pgClient       pg.Client
	chatRepository chatRepository.Repository
	chatService    chatService.Service
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			logger.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) PGClient(ctx context.Context) pg.Client {
	if s.pgClient == nil {
		pgCfg, err := pgxpool.ParseConfig(s.PGConfig().DSN())
		if err != nil {
			logger.Fatalf("failed to get db config: %s", err.Error())
		}

		client, err := pg.NewClient(ctx, pgCfg)
		if err != nil {
			logger.Fatalf("failed to initialize pg clients: %s", err.Error())
		}

		if err = client.PG().Ping(ctx); err != nil {
			logger.Fatalf("failed to ping pg: %s", err.Error())
		}

		closer.Add(client.Close)
		s.pgClient = client
	}
	return s.pgClient
}

func (s *serviceProvider) ChatRepository(ctx context.Context) chatRepository.Repository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.PGClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ServiceProvider(ctx context.Context) chatService.Service {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.ChatRepository(ctx))
	}

	return s.chatService
}

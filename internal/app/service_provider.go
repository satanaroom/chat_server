package app

import (
	"context"

	accessV1 "github.com/satanaroom/auth/pkg/access_v1"
	"github.com/satanaroom/auth/pkg/logger"
	"google.golang.org/grpc/credentials/insecure"

	authClient "github.com/satanaroom/chat_server/internal/clients/grpc/auth"
	"github.com/satanaroom/chat_server/internal/closer"
	"github.com/satanaroom/chat_server/internal/config"
	chatService "github.com/satanaroom/chat_server/internal/service/chat"
	"google.golang.org/grpc"
)

type serviceProvider struct {
	authConfig config.AuthClientConfig

	authClient  authClient.Client
	chatService chatService.Service
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) ChatService(ctx context.Context) chatService.Service {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.AuthClient(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) AuthClientConfig() config.AuthClientConfig {
	if s.authConfig == nil {
		cfg, err := config.NewAuthClientConfig()
		if err != nil {
			logger.Fatalf("failed to get auth client config: %s", err.Error())
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) AuthClient(ctx context.Context) authClient.Client {
	if s.authClient == nil {
		opts := grpc.WithTransportCredentials(insecure.NewCredentials())

		conn, err := grpc.DialContext(ctx, s.AuthClientConfig().Host(), opts)
		if err != nil {
			logger.Fatalf("failed to connect %s: %s", s.authConfig.Host(), err.Error())
		}
		closer.Add(conn.Close)

		client := accessV1.NewAccessV1Client(conn)
		s.authClient = authClient.NewClient(client)
	}

	return s.authClient
}

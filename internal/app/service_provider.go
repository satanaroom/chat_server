package app

import (
	"context"

	authV1 "github.com/satanaroom/auth/pkg/auth_v1"
	"github.com/satanaroom/auth/pkg/logger"

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

func (s *serviceProvider) AuthClient(_ context.Context) authClient.Client {
	if s.authClient == nil {
		conn, err := grpc.Dial(s.authConfig.Port(), grpc.WithDefaultCallOptions())
		if err != nil {
			logger.Fatalf("failed to connect %s: %s", s.authConfig.Port(), err.Error())
		}
		closer.Add(conn.Close)

		client := authV1.NewAuthV1Client(conn)
		s.authClient = authClient.NewClient(client)
	}

	return s.authClient
}

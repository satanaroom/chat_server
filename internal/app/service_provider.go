package app

import (
	"context"

	accessV1 "github.com/satanaroom/auth/pkg/access_v1"
	"github.com/satanaroom/auth/pkg/logger"
	chatV1 "github.com/satanaroom/chat_server/internal/api/chat_v1"
	"github.com/satanaroom/chat_server/internal/channels"
	authClient "github.com/satanaroom/chat_server/internal/clients/grpc/auth"
	"github.com/satanaroom/chat_server/internal/closer"
	"github.com/satanaroom/chat_server/internal/config"
	chatService "github.com/satanaroom/chat_server/internal/service/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type serviceProvider struct {
	authConfig    config.AuthClientConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
	tlsConfig     config.TLSConfig
	chatConfig    config.ChatConfig

	authClient  authClient.Client
	chatService chatService.Service

	tlsAuthCredentials   credentials.TransportCredentials
	tlsServerCredentials credentials.TransportCredentials

	chatImpl *chatV1.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) ChatService(ctx context.Context) chatService.Service {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.AuthClient(ctx), channels.NewChannels(s.ChatConfig().Capacity()))
	}

	return s.chatService
}

func (s *serviceProvider) AuthClientConfig() config.AuthClientConfig {
	if s.authConfig == nil {
		cfg, err := config.NewAuthClientConfig()
		if err != nil {
			logger.Fatalf("failed to get access client config: %s", err.Error())
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) AuthClient(ctx context.Context) authClient.Client {
	if s.authClient == nil {
		opts := grpc.WithTransportCredentials(s.TLSAuthCredentials(ctx))

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

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			logger.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			logger.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := config.NewSwaggerConfig()
		if err != nil {
			logger.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) TLSConfig() config.TLSConfig {
	if s.tlsConfig == nil {
		cfg, err := config.NewTLSConfig()
		if err != nil {
			logger.Fatalf("failed to get TLS config: %s", err.Error())
		}

		s.tlsConfig = cfg
	}

	return s.tlsConfig
}

func (s *serviceProvider) ChatConfig() config.ChatConfig {
	if s.chatConfig == nil {
		cfg, err := config.NewChatConfig()
		if err != nil {
			logger.Fatalf("failed to get chat config: %s", err.Error())
		}

		s.chatConfig = cfg
	}

	return s.chatConfig
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chatV1.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chatV1.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}

func (s *serviceProvider) TLSAuthCredentials(_ context.Context) credentials.TransportCredentials {
	if s.tlsAuthCredentials == nil {
		creds, err := credentials.NewClientTLSFromFile(s.TLSConfig().AuthCertFile(), "")
		if err != nil {
			logger.Fatalf("new auth client tls from file: %s", err.Error())
		}

		s.tlsAuthCredentials = creds
	}

	return s.tlsAuthCredentials
}

func (s *serviceProvider) TLSServerCredentials(_ context.Context) credentials.TransportCredentials {
	if s.tlsServerCredentials == nil {
		creds, err := credentials.NewServerTLSFromFile(s.TLSConfig().ServerCertFile(), s.TLSConfig().ServerKeyFile())
		if err != nil {
			logger.Fatalf("new server tls from file: %s", err.Error())
		}

		s.tlsServerCredentials = creds
	}

	return s.tlsServerCredentials
}

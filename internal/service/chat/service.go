package chat

import authClient "github.com/satanaroom/chat_server/internal/clients/grpc/auth"

var _ Service = (*service)(nil)

type Service interface {
}

type service struct {
	authClient authClient.Client
}

func NewService(authClient authClient.Client) *service {
	return &service{
		authClient: authClient,
	}
}

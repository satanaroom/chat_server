package chat

import (
	"context"

	"github.com/satanaroom/chat_server/internal/channels"
	authClient "github.com/satanaroom/chat_server/internal/clients/grpc/auth"
	"github.com/satanaroom/chat_server/internal/model"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

var _ Service = (*service)(nil)

type Service interface {
	CreateChat(ctx context.Context, usernames *model.CreateChat) (string, error)
	GetChatChannel(chatId string) (chan *chatV1.Message, error)
}

type service struct {
	authClient authClient.Client
	channels   channels.Channels
}

func NewService(authClient authClient.Client, channels channels.Channels) *service {
	return &service{
		authClient: authClient,
		channels:   channels,
	}
}

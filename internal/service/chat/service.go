package chat

import (
	"context"

	authClient "github.com/satanaroom/chat_server/internal/clients/grpc/auth"
	"github.com/satanaroom/chat_server/internal/model"
	"github.com/satanaroom/chat_server/internal/storage"
)

var _ Service = (*service)(nil)

type Service interface {
	CreateChat(ctx context.Context, usernames *model.CreateChat) (int64, error)
	ConnectChat(chatId int64) error
	SendMessage(message *model.Message)
	GetMessageChannel() <-chan *model.Message
}

type service struct {
	authClient  authClient.Client
	chatStorage storage.ChatStorage
}

func NewService(authClient authClient.Client, chatStorage storage.ChatStorage) *service {
	return &service{
		authClient:  authClient,
		chatStorage: chatStorage,
	}
}

package chat

import (
	"context"

	authClient "github.com/satanaroom/chat_server/internal/clients/grpc/auth"
	"github.com/satanaroom/chat_server/internal/model"
	"github.com/satanaroom/chat_server/internal/storage"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

var _ Service = (*service)(nil)

type Service interface {
	CreateChat(ctx context.Context, usernames *model.CreateChat) (string, error)
	ConnectChat(connectInfo *model.ConnectInfo, stream chatV1.ChatV1_ConnectChatServer) error
	SendMessage(chatId string, message *chatV1.Message) error
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

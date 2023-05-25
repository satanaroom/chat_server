package chat_v1

import (
	"context"

	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func (i *Implementation) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	i.chatService.CreateChat(ctx)

	return &desc.CreateChatResponse{
		ChatId: 1,
	}, nil
}

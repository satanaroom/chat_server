package chat_v1

import (
	"context"

	converter "github.com/satanaroom/chat_server/internal/converter/client"
	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	chatId, err := i.chatService.CreateChat(ctx, converter.ToCreateChatService(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "chatService.CreateChat: %s", err.Error())
	}

	return &desc.CreateChatResponse{
		ChatId: chatId,
	}, nil
}

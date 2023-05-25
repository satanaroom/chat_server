package chat_v1

import (
	"github.com/satanaroom/chat_server/internal/service/chat"
	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server

	chatService chat.Service
}

func NewImplementation(chatService chat.Service) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}

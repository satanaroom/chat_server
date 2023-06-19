package chat_v1

import (
	"sync"

	"github.com/satanaroom/chat_server/internal/service/chat"
	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server

	chatService chat.Service
	chats       Chats
}

type Chat struct {
	streams map[string]desc.ChatV1_ConnectChatServer
	mu      sync.RWMutex
}

type Chats struct {
	m  map[string]*Chat
	mu sync.RWMutex
}

func NewImplementation(chatService chat.Service) *Implementation {
	return &Implementation{
		chatService: chatService,
		chats:       Chats{m: make(map[string]*Chat)},
	}
}

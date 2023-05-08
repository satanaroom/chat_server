package chat

import (
	"github.com/satanaroom/chat_server/internal/repository/chat"
)

var _ Service = (*service)(nil)

type Service interface {
}

type service struct {
	chatRepository chat.Repository
}

func NewService(chatRepository chat.Repository) *service {
	return &service{
		chatRepository: chatRepository,
	}
}

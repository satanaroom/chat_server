package chat

import (
	"github.com/satanaroom/chat_server/internal/model"
)

func (s *service) SendMessage(message *model.Message) {
	s.chatStorage.SendMessage(message)
}

func (s *service) GetMessageChannel() <-chan *model.Message {
	return s.chatStorage.GetMessageChannel()
}

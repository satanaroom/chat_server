package chat

import (
	"github.com/satanaroom/chat_server/internal/sys"
	"github.com/satanaroom/chat_server/internal/sys/codes"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func (s *service) GetChatChannel(chatId string) (chan *chatV1.Message, error) {
	chatChannel, ok := s.channels.GetChannel(chatId)
	if !ok {
		return nil, sys.NewCommonError("channel not found", codes.NotFound)
	}

	return chatChannel, nil
}

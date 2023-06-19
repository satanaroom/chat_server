package chat

import (
	"github.com/satanaroom/chat_server/internal/errs"
	"github.com/satanaroom/chat_server/internal/sys"
	"github.com/satanaroom/chat_server/internal/sys/codes"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func (s *service) SendMessage(chatId string, message *chatV1.Message) error {
	chatChannel, ok := s.chatStorage.GetChannel(chatId)
	if !ok {
		return sys.NewCommonError(errs.ErrChannelNotFount.Error(), codes.NotFound)
	}

	chatChannel <- message

	return nil
}

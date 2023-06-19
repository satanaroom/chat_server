package chat

import (
	"github.com/satanaroom/chat_server/internal/errs"
	"github.com/satanaroom/chat_server/internal/model"
	"github.com/satanaroom/chat_server/internal/sys"
	"github.com/satanaroom/chat_server/internal/sys/codes"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func (s *service) ConnectChat(connectInfo *model.ConnectInfo, stream chatV1.ChatV1_ConnectChatServer) error {
	chatChannel, ok := s.chatStorage.GetChannel(connectInfo.ChatId)
	if !ok {
		return sys.NewCommonError(errs.ErrChannelNotFount.Error(), codes.NotFound)
	}
	s.chatStorage.CreateChat(connectInfo.ChatId)
	s.chatStorage.SetStream(connectInfo, stream)

	if err := s.chatStorage.HandleMessages(chatChannel, connectInfo, stream); err != nil {
		return sys.NewCommonError("failed to handle messages", codes.Internal)
	}

	return nil
}

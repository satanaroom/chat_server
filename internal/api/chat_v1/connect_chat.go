package chat_v1

import (
	converter "github.com/satanaroom/chat_server/internal/converter/server"
	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func (i *Implementation) ConnectChat(req *desc.ConnectChatRequest, stream desc.ChatV1_ConnectChatServer) error {
	if err := i.chatService.ConnectChat(converter.ToConnectInfo(req.GetChatId(), req.GetUsername()), stream); err != nil {
		return err
	}

	return nil
}

package chat_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func (i *Implementation) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	chatChannel, err := i.chatService.GetChatChannel(req.GetChatId())
	if err != nil {
		return nil, err
	}

	chatChannel <- req.GetMessage()

	return &empty.Empty{}, nil
}

package chat_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func (i *Implementation) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	if err := i.chatService.SendMessage(req.GetChatId(), req.GetMessage()); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

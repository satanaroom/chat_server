package chat_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	converter "github.com/satanaroom/chat_server/internal/converter/server"
	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func (i *Implementation) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	i.chatService.SendMessage(converter.ToMessage(req.GetMessage()))

	return &empty.Empty{}, nil
}

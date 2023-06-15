package server

import (
	"time"

	"github.com/satanaroom/chat_server/internal/model"
	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToMessage(message *desc.Message) *model.Message {
	return &model.Message{
		Text:   message.Text,
		From:   message.From,
		To:     message.To,
		SentAt: time.Now(),
	}
}

func ToChatUsers(chatInfo *model.CreateChat) *model.ChatUsers {
	return &model.ChatUsers{
		Usernames: chatInfo.Usernames,
	}
}

func ToConnectChatResponse(message *model.Message) *desc.ConnectChatResponse {
	return &desc.ConnectChatResponse{
		Message: &desc.Message{
			Text:   message.Text,
			From:   message.From,
			To:     message.To,
			SentAt: timestamppb.New(message.SentAt),
		},
	}
}

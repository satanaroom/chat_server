package client

import (
	"github.com/satanaroom/chat_server/internal/model"
	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func ToCreateChatService(req *desc.CreateChatRequest) *model.CreateChat {
	return &model.CreateChat{
		Usernames: req.GetUsernames(),
	}
}

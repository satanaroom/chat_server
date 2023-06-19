package server

import (
	"github.com/satanaroom/chat_server/internal/model"
)

func ToConnectInfo(chatId, username string) *model.ConnectInfo {
	return &model.ConnectInfo{
		ChatId:   chatId,
		Username: username,
	}
}

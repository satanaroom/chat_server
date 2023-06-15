package chat

import (
	"context"

	converter "github.com/satanaroom/chat_server/internal/converter/server"
	"github.com/satanaroom/chat_server/internal/model"
)

func (s *service) CreateChat(_ context.Context, usernames *model.CreateChat) (int64, error) {
	chatId := s.chatStorage.SaveChat(converter.ToChatUsers(usernames))

	return chatId, nil
}

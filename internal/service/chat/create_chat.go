package chat

import (
	"context"

	"github.com/google/uuid"
	"github.com/satanaroom/chat_server/internal/model"
	"github.com/satanaroom/chat_server/internal/sys"
	"github.com/satanaroom/chat_server/internal/sys/codes"
)

func (s *service) CreateChat(_ context.Context, usernames *model.CreateChat) (string, error) {
	chatID, err := uuid.NewUUID()
	if err != nil {
		return "", sys.NewCommonError("failed to generate UUID", codes.Internal)
	}

	s.channels.CreateChannel(chatID.String())

	return chatID.String(), nil
}

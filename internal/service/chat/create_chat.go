package chat

import (
	"context"

	"github.com/satanaroom/chat_server/internal/model"
)

func (s *service) CreateChat(ctx context.Context, usernames *model.CreateChat) (int64, error) {
	return 1, nil
}

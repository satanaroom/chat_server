package chat

import (
	"fmt"
)

func (s *service) ConnectChat(chatId int64) error {
	chat := s.chatStorage.GetChat(chatId)

	if chat != nil {
		return fmt.Errorf("chat is not exist")
	}

	return nil
}

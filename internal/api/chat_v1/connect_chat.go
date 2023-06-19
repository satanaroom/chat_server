package chat_v1

import (
	"github.com/satanaroom/chat_server/internal/sys"
	"github.com/satanaroom/chat_server/internal/sys/codes"
	desc "github.com/satanaroom/chat_server/pkg/chat_v1"
)

func (i *Implementation) ConnectChat(req *desc.ConnectChatRequest, stream desc.ChatV1_ConnectChatServer) error {
	chatId := req.GetChatId()
	username := req.GetUsername()

	chatChannel, err := i.chatService.GetChatChannel(chatId)
	if err != nil {
		return err
	}

	i.chats.mu.Lock()
	if _, ok := i.chats.m[chatId]; !ok {
		i.chats.m[chatId] = &Chat{
			streams: make(map[string]desc.ChatV1_ConnectChatServer),
		}
	}
	i.chats.mu.Unlock()

	i.chats.m[chatId].mu.Lock()
	i.chats.m[chatId].streams[username] = stream
	i.chats.m[chatId].mu.Unlock()

	for {
		select {
		case msg, ok := <-chatChannel:
			if !ok {
				return nil
			}

			for _, str := range i.chats.m[chatId].streams {
				if err = str.Send(&desc.ConnectChatResponse{
					Message: msg,
				}); err != nil {
					return sys.NewCommonError("failed to send message", codes.Internal)
				}
			}
		case <-stream.Context().Done():
			i.chats.m[chatId].mu.Lock()
			delete(i.chats.m[chatId].streams, username)
			i.chats.m[chatId].mu.Unlock()
			return nil
		}
	}
}

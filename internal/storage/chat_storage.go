package storage

import (
	"fmt"
	"sync"

	"github.com/satanaroom/chat_server/internal/model"
	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

type ChatStorage interface {
	CreateChannel(chatId string)
	GetChannel(chatId string) (chan *chatV1.Message, bool)
	CreateChat(chatId string)
	SetStream(connectInfo *model.ConnectInfo, stream chatV1.ChatV1_ConnectChatServer)
	HandleMessages(chatChannel chan *chatV1.Message, connectInfo *model.ConnectInfo, stream chatV1.ChatV1_ConnectChatServer) error
}

type Chat struct {
	streams map[string]chatV1.ChatV1_ConnectChatServer
	mu      sync.RWMutex
}

type Chats struct {
	m  map[string]*Chat
	mu sync.RWMutex
}

type Channels struct {
	m  map[string]chan *chatV1.Message
	mu sync.RWMutex
}

type chatStorage struct {
	chats    Chats
	channels Channels
	capacity int
}

func NewStorage(capacity int) *chatStorage {
	return &chatStorage{
		chats:    Chats{m: make(map[string]*Chat)},
		channels: Channels{m: make(map[string]chan *chatV1.Message)},
		capacity: capacity,
	}
}

func (c *chatStorage) CreateChannel(chatId string) {
	c.channels.m[chatId] = make(chan *chatV1.Message, 100)
}

func (c *chatStorage) GetChannel(chatId string) (chan *chatV1.Message, bool) {
	c.channels.mu.RLock()
	chatChannel, ok := c.channels.m[chatId]
	c.channels.mu.RUnlock()

	return chatChannel, ok
}

func (c *chatStorage) CreateChat(chatId string) {
	c.chats.mu.Lock()
	if _, ok := c.chats.m[chatId]; !ok {
		c.chats.m[chatId] = &Chat{
			streams: make(map[string]chatV1.ChatV1_ConnectChatServer),
		}
	}
	c.chats.mu.Unlock()
}

func (c *chatStorage) SetStream(connectInfo *model.ConnectInfo, stream chatV1.ChatV1_ConnectChatServer) {
	c.chats.m[connectInfo.ChatId].mu.Lock()
	c.chats.m[connectInfo.ChatId].streams[connectInfo.Username] = stream
	c.chats.m[connectInfo.ChatId].mu.Unlock()
}

func (c *chatStorage) HandleMessages(chatChannel chan *chatV1.Message, connectInfo *model.ConnectInfo, stream chatV1.ChatV1_ConnectChatServer) error {
	for {
		select {
		case msg, ok := <-chatChannel:
			if !ok {
				return nil
			}

			for _, str := range c.chats.m[connectInfo.ChatId].streams {
				if err := str.Send(&chatV1.ConnectChatResponse{
					Message: msg,
				}); err != nil {
					return fmt.Errorf("send message: %w", err)
				}
			}
		case <-stream.Context().Done():
			c.chats.m[connectInfo.ChatId].mu.Lock()
			delete(c.chats.m[connectInfo.ChatId].streams, connectInfo.Username)
			c.chats.m[connectInfo.ChatId].mu.Unlock()
			return nil
		}
	}
}

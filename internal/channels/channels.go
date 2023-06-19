package channels

import (
	"sync"

	chatV1 "github.com/satanaroom/chat_server/pkg/chat_v1"
)

type Channels interface {
	CreateChannel(chatId string)
	GetChannel(chatId string) (chan *chatV1.Message, bool)
}

type channels struct {
	m        map[string]chan *chatV1.Message
	mu       sync.RWMutex
	capacity int
}

func NewChannels(capacity int) *channels {
	return &channels{
		m:        make(map[string]chan *chatV1.Message),
		capacity: capacity,
	}
}

func (c *channels) CreateChannel(chatId string) {
	c.m[chatId] = make(chan *chatV1.Message, c.capacity)
}

func (c *channels) GetChannel(chatId string) (chan *chatV1.Message, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	chatChannel, ok := c.m[chatId]

	return chatChannel, ok
}

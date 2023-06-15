package storage

import (
	"sync"

	"github.com/satanaroom/chat_server/internal/model"
)

type ChatStorage interface {
	SaveChat(storage *model.ChatUsers) int64
	GetChat(chatId int64) *model.ChatUsers
	SendMessage(message *model.Message)
	GetMessageChannel() chan *model.Message
}

type chatStorage struct {
	messageCh chan *model.Message
	chatIds   int64
	storage   map[int64]*model.ChatUsers
	mu        *sync.Mutex
}

func NewStorage() *chatStorage {
	return &chatStorage{
		messageCh: make(chan *model.Message, 1),
		chatIds:   0,
		storage:   make(map[int64]*model.ChatUsers),
		mu:        &sync.Mutex{},
	}
}

func (c *chatStorage) SaveChat(storage *model.ChatUsers) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.chatIds++

	c.storage[c.chatIds] = storage

	return c.chatIds
}

func (c *chatStorage) GetChat(chatId int64) *model.ChatUsers {
	c.mu.Lock()
	defer c.mu.Unlock()

	storage, ok := c.storage[chatId]
	if !ok {
		return nil
	}

	return storage
}

func (c *chatStorage) SendMessage(message *model.Message) {
	c.messageCh <- message
}

func (c *chatStorage) GetMessageChannel() chan *model.Message {
	return c.messageCh
}

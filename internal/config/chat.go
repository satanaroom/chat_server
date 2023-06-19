package config

import (
	"github.com/satanaroom/auth/pkg/env"
)

var _ ChatConfig = (*chatConfig)(nil)

const chatCapacityEnvName = "GRPC_CAPACITY"

type ChatConfig interface {
	Capacity() int
}

type chatConfig struct {
	capacity int
}

func NewChatConfig() (*chatConfig, error) {
	var capacity int
	env.ToInt(&capacity, chatCapacityEnvName, 100)

	return &chatConfig{
		capacity: capacity,
	}, nil
}

func (c *chatConfig) Capacity() int {
	return c.capacity
}

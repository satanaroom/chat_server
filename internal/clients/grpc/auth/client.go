package auth

import (
	"context"
	"fmt"

	authV1 "github.com/satanaroom/auth/pkg/auth_v1"
	"github.com/satanaroom/chat_server/internal/converter/client/auth"
	"github.com/satanaroom/chat_server/internal/model"
)

var _ Client = (*client)(nil)

type Client interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
}

type client struct {
	authClient authV1.AuthV1Client
}

func NewClient(cl authV1.AuthV1Client) *client {
	return &client{
		authClient: cl,
	}
}

func (c *client) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	resp, err := c.authClient.Create(ctx, auth.ToCreateRequest(info))
	if err != nil {
		return 0, fmt.Errorf("authClient.Create: %w", err)
	}

	return auth.FromCreateResponse(resp), nil
}

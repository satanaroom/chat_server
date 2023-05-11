package auth

import (
	authV1 "github.com/satanaroom/auth/pkg/auth_v1"

	"github.com/satanaroom/chat_server/internal/model"
)

func ToCreateRequest(info *model.UserInfo) *authV1.CreateRequest {
	return &authV1.CreateRequest{
		Info: &authV1.UserInfo{
			Username: info.User.Username,
			Email:    info.User.Email,
			Password: info.User.Password,
			Role:     authV1.Role(info.User.Role),
		},
		PasswordConfirm: info.PasswordConfirm,
	}
}

func FromCreateResponse(resp *authV1.CreateResponse) int64 {
	return resp.GetId()
}

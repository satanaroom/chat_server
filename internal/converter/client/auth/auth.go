package auth

//func ToCreateRequest(info *model.UserInfo) *accessV1.CreateRequest {
//	return &accessV1.CreateRequest{
//		Info: &authV1.UserInfo{
//			Username: info.User.Username,
//			Email:    info.User.Email,
//			Password: info.User.Password,
//			Role:     authV1.Role(info.User.Role),
//		},
//		PasswordConfirm: info.PasswordConfirm,
//	}
//}

//func FromCreateResponse(resp *accessV1.CreateResponse) int64 {
//	return resp.GetId()
//}

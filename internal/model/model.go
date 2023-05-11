package model

type Role int64

const (
	RoleAdmin Role = 1
	RoleUser  Role = 2
)

type User struct {
	Username string
	Email    string
	Password string
	Role     Role
}

type UserInfo struct {
	User            User
	PasswordConfirm string
}

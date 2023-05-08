package errs

import "errors"

var ErrUsernameEmpty = errors.New("username is empty")
var ErrPasswordEmpty = errors.New("password is empty")
var ErrPasswordMismatch = errors.New("password and password confirm do not match")
var ErrRoleInvalid = errors.New("role is invalid")
var ErrEmailInvalid = errors.New("email is invalid")
var ErrUserNotFound = errors.New("user not found")
var ErrDSNNotFound = errors.New("pg dsn not found")
var ErrGRPCPortNotFound = errors.New("grpc port not found")

package user

import "errors"

// 用errors包定义的错误处理
var (
	ErrUserExisWithName    = errors.New("用户已经存在")
	ErrUserNotFound        = errors.New("用户没有找到")
	ErrMismatchedPasswords = errors.New("密码错误")
	ErrInvalidUsername     = errors.New("无效的用户名")
	ErrInvalidPassword     = errors.New("无效的密码")
)

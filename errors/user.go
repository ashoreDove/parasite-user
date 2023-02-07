package errors

import "errors"

var (
	CodeErr        = errors.New("验证码不正确")
	UserExistErr   = errors.New("账户已经存在")
	PwdErr         = errors.New("密码错误")
	UserNoExistErr = errors.New("用户不存在")
)

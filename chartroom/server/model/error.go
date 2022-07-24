package model

import "errors"

//自定义错误
var (
	Err_USER_NOTEXISTS = errors.New("用户不存在")
	Err_USER_EXISTS    = errors.New("用户已注册")
	Err_USER_PWD       = errors.New("密码错误")
)

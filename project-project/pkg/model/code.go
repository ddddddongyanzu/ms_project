package model

import (
	"test.com/project-common/errs"
)

var (
	RedisError           = errs.NewError(999, "redis 错误")
	DBError              = errs.NewError(998, "DB 错误")
	NoLegalMobile        = errs.NewError(10102001, "手机号不合法")
	CaptchaNotExistError = errs.NewError(10102002, "验证码不存在或已过期")
	CaptchaError         = errs.NewError(10102003, "验证码错误")
	EmailExist           = errs.NewError(10102004, "邮箱已存在")
	AccountExist         = errs.NewError(10102005, "账户已存在")
	MobileExist          = errs.NewError(10102006, "手机号已存在")
	AccountAndPwdError   = errs.NewError(10102007, "账号密码不正确")
)

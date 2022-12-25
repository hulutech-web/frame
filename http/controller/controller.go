package controller

import (
	"gitee.com/hulutech/frame/policy"
	"gitee.com/hulutech/frame/request/http/auth"
	"gitee.com/hulutech/frame/validator"
)

type Controller interface {
	Validate(c validator.Context, _validator interface{}, onlyFirstError bool) (isAbort bool)

	Authorize(c policy.Context, policies policy.Policier, action policy.Action) (permit bool, user auth.IUser)
}

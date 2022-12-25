package controller

import (
	"github.com/hulutech-web/frame/policy"
	"github.com/hulutech-web/frame/request/http/auth"
	"github.com/hulutech-web/frame/validator"
)

type Controller interface {
	Validate(c validator.Context, _validator interface{}, onlyFirstError bool) (isAbort bool)

	Authorize(c policy.Context, policies policy.Policier, action policy.Action) (permit bool, user auth.IUser)
}

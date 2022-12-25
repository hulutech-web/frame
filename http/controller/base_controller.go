package controller

import (
	"gitee.com/hulutech/frame/policy"
	"gitee.com/hulutech/frame/request/http/auth"
	"gitee.com/hulutech/frame/validator"
)

type BaseController struct {
	policy.Authorization
	auth.RequestUser
	validator.Validation
}

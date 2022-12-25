package controller

import (
	"github.com/hulutech-web/frame/policy"
	"github.com/hulutech-web/frame/request/http/auth"
	"github.com/hulutech-web/frame/validator"
)

type BaseController struct {
	policy.Authorization
	auth.RequestUser
	validator.Validation
}

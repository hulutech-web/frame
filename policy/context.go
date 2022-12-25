package policy

import (
	"gitee.com/hulutech/frame/context"
	"gitee.com/hulutech/frame/request/http/auth"
)

type Context interface {
	context.LifeCycleContextor
	context.ResponseContextor
	auth.Context
	auth.RequestIUser
}

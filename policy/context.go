package policy

import (
	"github.com/hulutech-web/frame/context"
	"github.com/hulutech-web/frame/request/http/auth"
)

type Context interface {
	context.LifeCycleContextor
	context.ResponseContextor
	auth.Context
	auth.RequestIUser
}

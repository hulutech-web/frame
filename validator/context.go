package validator

import (
	"gitee.com/hulutech/frame/context"
	"gitee.com/hulutech/frame/resources/lang"
)

type Context interface {
	context.RequestBindingContextor
	context.ResponseContextor
	lang.Context
}

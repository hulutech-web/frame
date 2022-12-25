package validator

import (
	"github.com/hulutech-web/frame/context"
	"github.com/hulutech-web/frame/resources/lang"
)

type Context interface {
	context.RequestBindingContextor
	context.ResponseContextor
	lang.Context
}

package middleware

import (
	"github.com/hulutech-web/frame/policy"
	"github.com/hulutech-web/frame/request"
)

func Policy(_policy policy.Policier, action policy.Action) request.HandlerFunc {
	return func(c request.Context) {
		policy.Middleware(_policy, action, c, c.Params())
	}
}

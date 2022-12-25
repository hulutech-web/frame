package middleware

import (
	"gitee.com/hulutech/frame/policy"
	"gitee.com/hulutech/frame/request"
)

func Policy(_policy policy.Policier, action policy.Action) request.HandlerFunc {
	return func(c request.Context) {
		policy.Middleware(_policy, action, c, c.Params())
	}
}

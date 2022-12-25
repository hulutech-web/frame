package middleware

import (
	"gitee.com/hulutech/frame/request"
	"gitee.com/hulutech/frame/request/http/auth"
)

func IUser(userModelPtr auth.IUser) request.HandlerFunc {
	return func(c request.Context) {
		c.SetIUserModel(userModelPtr)

		c.Next()
	}
}

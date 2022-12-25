package middleware

import (
	"github.com/hulutech-web/frame/request"
	"github.com/hulutech-web/frame/request/http/auth"
)

func IUser(userModelPtr auth.IUser) request.HandlerFunc {
	return func(c request.Context) {
		c.SetIUserModel(userModelPtr)

		c.Next()
	}
}

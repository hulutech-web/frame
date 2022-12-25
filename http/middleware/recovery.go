package middleware

import (
	"github.com/gin-gonic/gin"

	"gitee.com/hulutech/frame/request"
)

func Recovery() request.HandlerFunc {
	return func(c request.Context) {
		gin.Recovery()(c.GinContext())
	}
}

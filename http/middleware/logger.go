package middleware

import (
	"github.com/gin-gonic/gin"

	"gitee.com/hulutech/frame/request"
)

func Logger() request.HandlerFunc {
	return func(c request.Context) {
		gin.Logger()(c.GinContext())
	}
}

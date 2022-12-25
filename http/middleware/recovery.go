package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/hulutech-web/frame/request"
)

func Recovery() request.HandlerFunc {
	return func(c request.Context) {
		gin.Recovery()(c.GinContext())
	}
}

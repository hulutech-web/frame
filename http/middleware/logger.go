package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/hulutech-web/frame/request"
)

func Logger() request.HandlerFunc {
	return func(c request.Context) {
		gin.Logger()(c.GinContext())
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/hulutech-web/frame/request"
)

func BasicAuth(accounts map[string]string) request.HandlerFunc {
	return func(c request.Context) {
		gin.BasicAuth(accounts)(c.GinContext())
	}
}

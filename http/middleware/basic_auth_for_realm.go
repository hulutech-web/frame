package middleware

import (
	"github.com/gin-gonic/gin"

	"gitee.com/hulutech/frame/request"
)

func BasicAuthWithRealm(accounts map[string]string, realm string) request.HandlerFunc {
	return func(c request.Context) {
		gin.BasicAuthForRealm(accounts, realm)(c.GinContext())
	}
}

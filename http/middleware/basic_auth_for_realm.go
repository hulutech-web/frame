package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/hulutech-web/frame/request"
)

func BasicAuthWithRealm(accounts map[string]string, realm string) request.HandlerFunc {
	return func(c request.Context) {
		gin.BasicAuthForRealm(accounts, realm)(c.GinContext())
	}
}

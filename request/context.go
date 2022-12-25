package request

import (
	"github.com/gin-gonic/gin"

	"github.com/hulutech-web/frame/context"
	"github.com/hulutech-web/frame/request/http/auth"
	"github.com/hulutech-web/frame/utils/jwt"
)

type Context interface {
	// http context
	context.HttpContextor

	GinContext() *gin.Context

	SetAuthClaim(claims *jwt.UserClaims) //@todo move into a new interface
	SetIUserModel(iUser auth.IUser)      //@todo move into a new interface

	auth.Context
	auth.RequestIUser
}

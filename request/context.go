package request

import (
	"github.com/gin-gonic/gin"

	"gitee.com/hulutech/frame/context"
	"gitee.com/hulutech/frame/request/http/auth"
	"gitee.com/hulutech/frame/utils/jwt"
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

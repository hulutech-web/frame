package middleware

import (
	l "gitee.com/hulutech/frame/helpers/locale"
	"gitee.com/hulutech/frame/request"

	"gitee.com/hulutech/frame/config"
)

func Locale() request.HandlerFunc {
	return func(c request.Context) {
		locale := c.Request().Header.Get("locale")
		if locale == "" {
			locale = c.DefaultQuery("locale", config.GetString("app.locale"))
		}

		l.SetLocale(c, locale)

		c.Next()
	}
}

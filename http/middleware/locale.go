package middleware

import (
	l "github.com/hulutech-web/frame/helpers/locale"
	"github.com/hulutech-web/frame/request"

	"github.com/hulutech-web/frame/config"
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

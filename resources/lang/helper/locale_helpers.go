package helper

import (
	"gitee.com/hulutech/frame/resources/lang"

	"gitee.com/hulutech/frame/config"

	"gitee.com/hulutech/frame/resources/lang/internal"
)

func AddLocale(langName string, customTranslation *lang.CustomTranslation, validationTranslation *lang.ValidationTranslation) {
	internal.AddLocale(langName, customTranslation, validationTranslation)
}

func SetLocale(c lang.Context, langName string) {
	c.Set("locale", langName)
}
func Locale(c lang.Context) string {
	if contextLocale, exist := c.Get("locale"); exist {
		l := contextLocale.(string)
		return fallbackLocale(l)
	}
	configLocale := config.GetString("app.locale")
	return fallbackLocale(configLocale)
}

func fallbackLocale(langName string) string {
	if !internal.HasLocale(langName) {
		return config.GetString("app.fallback_locale", "en")
	}
	return langName
}

package trans

import (
	"github.com/hulutech-web/frame/resources/lang"
	"gopkg.in/go-playground/validator.v9"

	"github.com/hulutech-web/frame/resources/lang/helper"
)

func ValidationTranslate(v *validator.Validate, langName string, e validator.ValidationErrors) lang.ValidationError {
	return helper.ValidationTranslate(v, langName, e)
}
func CustomTranslate(messageID string, data map[string]interface{}, langName string) string {
	return helper.CustomTranslate(messageID, data, langName)
}

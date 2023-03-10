package helper

import (
	"errors"

	"github.com/hulutech-web/frame/resources/lang"
	"github.com/hulutech-web/frame/resources/lang/internal"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

func ValidationTranslate(v *validator.Validate, langName string, e validator.ValidationErrors) lang.ValidationError {
	t, err := translator(v, langName)
	if err != nil {
		return lang.ValidationError{
			"error": err.Error(),
		}
	}
	return lang.ValidationError(e.Translate(t))
}

func translator(v *validator.Validate, langName string) (ut.Translator, error) {

	l := internal.Locale(langName)
	if l == nil {
		return nil, errors.New("validation translation not found")
	}

	if !l.ValidationRegistered() {
		if err := internal.RegisterDefaultTranslations(v, l); err != nil {
			return nil, err
		}
	}

	return l.UniversalTranslator(), nil
}

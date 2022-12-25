package helpers

import (
	"github.com/hulutech-web/frame/helpers/trans"
	"github.com/hulutech-web/frame/request"

	"github.com/hulutech-web/frame/helpers/locale"
)

func L(c request.Context, messageID string, dataNlocale ...interface{}) string {
	l := locale.Locale(c)
	data := make(map[string]interface{})
	switch len(dataNlocale) {
	case 1:
		data = dataNlocale[0].(map[string]interface{})
		break
	case 2:
		l = dataNlocale[1].(string)
		break
	default:
	}

	return trans.CustomTranslate(messageID, data, l)
}

//func Encrypt(secret string) {
//	//@todo
//}
//
//func Decrypt() {
//	//@todo
//}

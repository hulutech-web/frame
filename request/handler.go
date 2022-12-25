package request

import (
	"github.com/gin-gonic/gin"

	"gitee.com/hulutech/frame/request/http"
)

func ConvertHandlers(handlers []HandlerFunc) (ginHandlers []gin.HandlerFunc) {
	for _, h := range handlers {
		handler := h

		ginHandlers = append(ginHandlers, func(c *gin.Context) {
			tmaicContext := http.ConvertContext(c)
			handler(tmaicContext)
		})
	}
	return
}

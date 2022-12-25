package views

import (
	"github.com/hulutech-web/frame/request"
	"github.com/hulutech-web/frame/view"
)

func Initialize(r *request.Engine) {
	view.Initialize(r)
}

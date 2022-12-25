package routes

import (
	"github.com/hulutech-web/frame/monitor/routes/versions"
	"github.com/hulutech-web/frame/request"
	"github.com/hulutech-web/frame/route"
)

func Register(router *request.Engine) {
	defer route.Bind(router)

	versions.NewMonitor(router)
}

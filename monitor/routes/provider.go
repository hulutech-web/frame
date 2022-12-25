package routes

import (
	"gitee.com/hulutech/frame/monitor/routes/versions"
	"gitee.com/hulutech/frame/request"
	"gitee.com/hulutech/frame/route"
)

func Register(router *request.Engine) {
	defer route.Bind(router)

	versions.NewMonitor(router)
}

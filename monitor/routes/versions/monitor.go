package versions

import (
	"gitee.com/hulutech/frame/config"
	"gitee.com/hulutech/frame/http/middleware"
	"gitee.com/hulutech/frame/monitor/routes/groups"

	"gitee.com/hulutech/frame/request"
	"gitee.com/hulutech/frame/route"
)

func NewMonitor(engine *request.Engine) {
	ver := route.NewVersion(engine, "monitor")

	accounts := make(map[string]string)
	accounts[config.GetString("monitor.username")] = config.GetString("monitor.password")

	// noauth routes
	ver.NoAuth("", func(grp route.Grouper) {
		grp.AddGroup("/dashboard", &groups.DashboardGroup{})
	}, middleware.BasicAuth(accounts))
}

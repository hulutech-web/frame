package versions

import (
	"github.com/hulutech-web/frame/config"
	"github.com/hulutech-web/frame/http/middleware"
	"github.com/hulutech-web/frame/monitor/routes/groups"

	"github.com/hulutech-web/frame/request"
	"github.com/hulutech-web/frame/route"
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

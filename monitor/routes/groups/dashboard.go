package groups

import (
	"gitee.com/hulutech/frame/monitor/app/http/controllers"

	"gitee.com/hulutech/frame/route"
)

type DashboardGroup struct {
	DashboardController          controllers.Dashboard
	DashboardWebsocketController controllers.DashboardWebsocketController
}

func (dg *DashboardGroup) Group(group route.Grouper) {
	group.GET("/", dg.DashboardController.Index)
	group.Websocket("/ws", &dg.DashboardWebsocketController)
}

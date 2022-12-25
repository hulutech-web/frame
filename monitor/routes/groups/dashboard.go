package groups

import (
	"github.com/hulutech-web/frame/monitor/app/http/controllers"

	"github.com/hulutech-web/frame/route"
)

type DashboardGroup struct {
	DashboardController          controllers.Dashboard
	DashboardWebsocketController controllers.DashboardWebsocketController
}

func (dg *DashboardGroup) Group(group route.Grouper) {
	group.GET("/", dg.DashboardController.Index)
	group.Websocket("/ws", &dg.DashboardWebsocketController)
}

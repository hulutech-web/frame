package controllers

import (
	"net/http"

	"github.com/hulutech-web/frame/config"
	"github.com/hulutech-web/frame/helpers/toto"
	"github.com/hulutech-web/frame/http/controller"
	"github.com/hulutech-web/frame/request"
)

type Dashboard struct {
	controller.BaseController
}

func (d *Dashboard) Index(c request.Context) {
	c.HTML(http.StatusOK, "hulu.index", toto.V{
		"url": "ws://" + ":" + config.GetString("monitor.port") + "/monitor/dashboard/ws",
	})
	return
}

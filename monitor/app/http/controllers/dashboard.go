package controllers

import (
	"net/http"

	"gitee.com/hulutech/frame/config"
	"gitee.com/hulutech/frame/helpers/toto"
	"gitee.com/hulutech/frame/http/controller"
	"gitee.com/hulutech/frame/request"
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

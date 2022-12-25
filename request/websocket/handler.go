package websocket

import (
	"net/http"

	"gitee.com/hulutech/frame/helpers/zone"
	"gitee.com/hulutech/frame/http/controller"
	"gitee.com/hulutech/frame/request"
)

type Handler interface {
	DefaultChannels() []string
	OnMessage(hub Hub, msg *Msg)
	Loop(hub Hub) error

	OnPing(hub Hub, appData string)
	OnPong(hub Hub, appData string)
	OnClose(hub Hub, code int, text string)

	configer
	controller.Controller
}
type Hub interface {
	Send(msg *Msg)
	Broadcast(msg *Msg)
	BroadcastTo(channelName string, msg *Msg)

	name() string
	available() bool

	channeller
	request.Context
}
type configer interface {
	ReadBufferSize() int
	WriteBufferSize() int
	CheckOrigin(r *http.Request) bool
	WriteTimeout() zone.Duration
	ReadTimeout() zone.Duration
	MaxMessageSize() int64
}

type channeller interface {
	JoinChannel(channelName string)
	LeaveChannel(channelName string)
}

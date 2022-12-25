package driver

import (
	"time"
	"github.com/golang/protobuf/proto"

	//"gitee.com/hulutech/frame/helpers/zone"
)

type ProtoCacheGetter interface {
	Pget(key string, valuePtr proto.Message, defaultValuePtr ...proto.Message) error
}
type ProtoCacher interface {
	Ppull(key string, valuePtr proto.Message, defaultValuePtr ...proto.Message) error
	Pput(key string, valuePtr proto.Message, future time.Duration) bool
	Padd(key string, valuePtr proto.Message, future time.Duration) bool
	Pforever(key string, valuePtr proto.Message) bool

	ProtoCacheGetter
}

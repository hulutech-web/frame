package cache

import (
	"gitee.com/hulutech/frame/cache/driver"
)

type cacher interface {
	driver.ProtoCacher
	driver.BasicCacher
}

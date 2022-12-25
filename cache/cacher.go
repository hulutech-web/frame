package cache

import (
	"github.com/hulutech-web/frame/cache/driver"
)

type cacher interface {
	driver.ProtoCacher
	driver.BasicCacher
}

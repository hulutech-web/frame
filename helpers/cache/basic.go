package cache

import (
	c "gitee.com/hulutech/frame/cache"
	//"gitee.com/hulutech/frame/helpers/zone"
	"gitee.com/hulutech/frame/helpers/log"
	"time"
)

func Prefix() string {
	return c.Cache().Prefix()
}

func Has(key string) bool {
	return c.Cache().Has(key)
}

 

func Get(key string, defaultValue ...interface{}) interface{} {
	return c.Cache().Get(key, defaultValue...)
}
func Pull(key string, defaultValue ...interface{}) interface{} {
	return c.Cache().Pull(key, defaultValue...)
}
func Put(key string, value interface{}, future  time.Duration) bool {
	log.Debug(future)

	return c.Cache().Put(key, value, future)
}
func Add(key string, value interface{}, future time.Duration) bool {
	return c.Cache().Add(key, value, future)
}
func Increment(key string, value int64) (incremented int64, success bool) {
	return c.Cache().Increment(key, value)
}
func Decrement(key string, value int64) (decremented int64, success bool) {
	return c.Cache().Decrement(key, value)
}
func Forever(key string, value interface{}) bool {
	return c.Cache().Forever(key, value)
}
func Forget(key string) bool {
	return c.Cache().Forget(key)
}

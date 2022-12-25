package monitor

import "gitee.com/hulutech/frame/monitor/app/logics/dashboard"

func Shutdown() error {
	dashboard.Flow.Close()
	return nil
}

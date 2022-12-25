package monitor

import "github.com/hulutech-web/frame/monitor/app/logics/dashboard"

func Shutdown() error {
	dashboard.Flow.Close()
	return nil
}

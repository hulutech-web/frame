package monitor

import (
	"context"
	"net/http"
	"sync"

	"gitee.com/hulutech/frame/monitor/resources/views"

	c "gitee.com/hulutech/frame/config"
	"gitee.com/hulutech/frame/helpers/log"
	"gitee.com/hulutech/frame/helpers/toto"
	"gitee.com/hulutech/frame/helpers/zone"
	"gitee.com/hulutech/frame/monitor/routes"
	"gitee.com/hulutech/frame/request"
)

func HttpMonitorServe(parentCtx context.Context, wg *sync.WaitGroup) {
	r := request.New()

	//sentry.Use(r.GinEngine(), false)

	//r.Use(middleware.IUser(&models.YourUserModel{})) // set default auth user model, or use config auth.model_ptr

	routes.Register(r)

	views.Initialize(r)

	s := &http.Server{
		Addr:           ":" + c.GetString("monitor.port"),
		Handler:        r,
		ReadTimeout:    zone.Duration(c.GetInt64("app.read_timeout_seconds")) * zone.Second,
		WriteTimeout:   zone.Duration(c.GetInt64("app.write_timeout_seconds")) * zone.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Info("Monitor Served At", toto.V{"Addr": s.Addr})
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	<-parentCtx.Done()

	log.Info("Shutdown Monitor Server ...")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	ctx, cancel := context.WithTimeout(parentCtx, 5*zone.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Monitor Server Shutdown: ", toto.V{"error": err})
	}

	wg.Done()
}

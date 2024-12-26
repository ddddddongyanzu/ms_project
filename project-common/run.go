package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run(r *gin.Engine, srvName string, addr string, stop func()) {

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Printf("Shutdown Server %s...\n", srvName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if stop != nil {
		stop()
	}
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s Server Shutdown: %v\n", srvName, err)
	}
	log.Printf("%s Server exiting\n", srvName)
}

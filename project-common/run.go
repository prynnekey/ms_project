package common

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine, addr string, srvName string) {

	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 保证下面的优雅启停
	go func() {
		log.Printf("%s is running at %s\n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 2 秒的超时时间）
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT (Ctrl+C触发)
	// kill -9 is syscall.SIGKILL but can"t be catch, so don"t need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞直至有信号传入
	<-quit
	log.Printf("Shutting Down %s ...\n", srvName)

	// 创建一个2秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 优雅地关闭服务器，等待所有连接都关闭
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s Shutdown: %v", srvName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("Wait timeout...")
	}
	log.Printf("%s stop success.\n", srvName)
}

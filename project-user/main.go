package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/prynnekey/ms_project/project-common"
	"github.com/prynnekey/ms_project/project-user/config"
	"github.com/prynnekey/ms_project/project-user/router"
)

func main() {
	r := gin.Default()
	// 路由
	router.Init(r)

	// grpc服务注册
	grpc := router.RegisterGrpc()
	stop := func() {
		grpc.Stop()
	}

	srv.Run(r, config.AppConfig.ServerConfig.Addr, config.AppConfig.ServerConfig.Name, stop)
}

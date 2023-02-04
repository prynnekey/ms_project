package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/prynnekey/ms_project/project-api/api"
	"github.com/prynnekey/ms_project/project-api/config"
	"github.com/prynnekey/ms_project/project-api/router"
	srv "github.com/prynnekey/ms_project/project-common"
)

func main() {
	r := gin.Default()
	// 路由
	router.Init(r)

	srv.Run(r, config.AppConfig.SC.Addr, config.AppConfig.SC.Name, nil)
}

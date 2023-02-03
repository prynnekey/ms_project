package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/prynnekey/ms_project/project-common"
	_ "github.com/prynnekey/ms_project/project-user/api"
	"github.com/prynnekey/ms_project/project-user/router"
)

func main() {
	r := gin.Default()

	router.Init(r)

	srv.Run(r, ":80", "project user")
}

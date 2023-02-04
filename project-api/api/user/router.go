package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/ms_project/project-api/router"
)

type RouterUser struct {
}

func init() {
	log.Println("init user router")
	router.Register(&RouterUser{})
}

func (*RouterUser) Route(r *gin.Engine) {
	// 初始化grpc链接
	InitRpcUserClient()

	h := New()
	r.POST("project/login/getCaptcha", h.GetCaptcha())
}

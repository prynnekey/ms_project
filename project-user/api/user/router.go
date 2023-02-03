package user

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/ms_project/project-user/router"
)

func init() {
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}
	r.GET("/user", h.Hello())
}

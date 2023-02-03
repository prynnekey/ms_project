package user

import "github.com/gin-gonic/gin"

type HandlerUser struct {
}

func (*HandlerUser) Hello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, "hello")
	}
}

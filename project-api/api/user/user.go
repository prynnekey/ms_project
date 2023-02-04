package user

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	common "github.com/prynnekey/ms_project/project-common"
	"github.com/prynnekey/ms_project/project-common/errs"
	login_service_v1 "github.com/prynnekey/ms_project/project-user/pkg/service/login.service.v1"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (*HandlerUser) GetCaptcha() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := common.Result{}
		// 获取参数
		mobile := ctx.PostForm("mobile")

		// 调用gprc中的代码
		c, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		cr, err := LoginServiceClient.Captcha(c, &login_service_v1.CaptchaRequest{Mobile: mobile})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			ctx.JSON(200, res.Fail(code, msg))
			return
		}

		ctx.JSON(200, res.Success(cr))
	}
}

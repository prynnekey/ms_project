package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	common "github.com/prynnekey/ms_project/project-common"
	"github.com/prynnekey/ms_project/project-user/pkg/dao"
	"github.com/prynnekey/ms_project/project-user/pkg/model"
	"github.com/prynnekey/ms_project/project-user/pkg/repo"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{
		cache: dao.Rc,
	}
}

// 获取验证码
func (h *HandlerUser) GetCaptcha() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := common.Result{}
		// 1.获取参数
		mobile := ctx.PostForm("mobile")
		// 2.验证参数
		if !common.ValidateMobile(mobile) {
			ctx.JSON(http.StatusOK, res.Fail(model.IllegalMobile, "手机号码格式不正确"))
			return
		}
		// 3.生成验证码
		code := common.GenerateCaptcha(6)
		// 4.发送验证码(放入go协程中,接口可以快速响应)
		go func() {
			// 调用短信服务发送验证码
			time.Sleep(time.Second * 2)
			fmt.Println("验证码发送成功", code)
			// 5.将验证码保存到redis中
			c, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()
			err := h.cache.Put(c, "REGISTER_"+mobile, code, time.Minute*5)
			if err != nil {
				fmt.Println("redis保存验证码失败", err)
			}
		}()

		// 6.返回结果
		ctx.JSON(http.StatusOK, res.Success("发送成功"))
	}
}

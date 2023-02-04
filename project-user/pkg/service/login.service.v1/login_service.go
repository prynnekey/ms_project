package login_service_v1

import (
	context "context"
	"fmt"
	"time"

	common "github.com/prynnekey/ms_project/project-common"
	"github.com/prynnekey/ms_project/project-common/errs"
	"github.com/prynnekey/ms_project/project-user/pkg/dao"
	"github.com/prynnekey/ms_project/project-user/pkg/model"
	"github.com/prynnekey/ms_project/project-user/pkg/repo"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (ls *LoginService) Captcha(ctx context.Context, req *CaptchaRequest) (*CaptchaResponse, error) {
	// 1.获取参数
	mobile := req.Mobile
	// 2.验证参数
	if !common.ValidateMobile(mobile) {
		return nil, errs.GrpcError(model.IllegalMobile)
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
		err := ls.cache.Put(c, "REGISTER_"+mobile, code, time.Minute*5)
		if err != nil {
			fmt.Println("redis保存验证码失败", err)
		}
	}()

	// 6.返回结果
	return &CaptchaResponse{Code: 200, Message: "验证码发送成功"}, nil
}

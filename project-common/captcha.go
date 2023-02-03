package common

import (
	"math/rand"
	"strconv"
)

// 生成验证码
func GenerateCaptcha(len int) string {
	captcha := ""
	for i := 0; i < len; i++ {
		captcha += strconv.Itoa(rand.Intn(10))
	}
	return captcha
}

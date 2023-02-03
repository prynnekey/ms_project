package common

import "regexp"

// 验证手机号是否合法
func ValidateMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	regular := "^((13[0-9])|(14[5,7,9])|(15([0-3]|[5-9]))|(16[6])|(17[0,1,3,5,6,7,8])|(18[0-9])|(19[8,9]))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

package model

import (
	"github.com/prynnekey/ms_project/project-common/errs"
)

var (
	IllegalMobile = errs.New(2001, "手机号不合法") // 非法手机号
)

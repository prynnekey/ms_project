package common

// Result 返回结果
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewResult 返回结果
func NewResult(code int, msg string, data interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// NewSuccessResult 返回成功结果
func (*Result) Success(data interface{}) *Result {
	return NewResult(200, "success", data)
}

// NewErrorResult 返回错误结果
func (*Result) Fail(code int, msg string) *Result {
	return NewResult(code, msg, nil)
}

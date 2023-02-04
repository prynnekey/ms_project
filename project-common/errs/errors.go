package errs

import "fmt"

type ErrorCode int

type BError struct {
	Code    ErrorCode
	Message string
}

func (e BError) Error() string {
	return fmt.Sprintf("code:%v,msg:%s", e.Code, e.Message)
}

func New(code ErrorCode, message string) *BError {
	return &BError{code, message}
}

package errs

import (
	common "github.com/prynnekey/ms_project/project-common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcError(err *BError) error {
	return status.Error(codes.Code(err.Code), err.Message)
}

func ParseGrpcError(err error) (code common.BusinessCode, msg string) {
	fromError, _ := status.FromError(err)
	return common.BusinessCode(fromError.Code()), fromError.Message()
}

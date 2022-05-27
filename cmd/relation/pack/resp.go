package pack

import (
	"errors"
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResponse(err error) *relation.BaseResponse {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *relation.BaseResponse {
	return &relation.BaseResponse{
		StatusCode:  err.ErrCode,
		StatusMsg:   err.ErrMsg,
		ServiceTime: time.Now().Unix(),
	}
}

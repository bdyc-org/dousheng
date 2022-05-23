package pack

import (
	"errors"
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

func BuildBaseResponse(err error) *user.BaseResponse {
	if err == nil {
		return baseResponse(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResponse(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResponse(s)
}

func baseResponse(err errno.ErrNo) *user.BaseResponse {
	return &user.BaseResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg, ServiceTime: time.Now().Unix()}
}

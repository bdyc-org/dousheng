package pack

import (
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"time"
)

func BuildBaseResponse(statusCode int64, statusMsg string) *comment.BaseResponse {
	return &comment.BaseResponse{
		StatusCode:  statusCode,
		StatusMsg:   statusMsg,
		ServiceTime: time.Now().Unix(),
	}
}

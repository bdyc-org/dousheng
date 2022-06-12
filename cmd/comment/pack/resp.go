package pack

import (
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/comment"
)

func BuildBaseResponse(statusCode int64, statusMsg string) *comment.BaseResponse {
	return &comment.BaseResponse{
		StatusCode:  statusCode,
		StatusMsg:   statusMsg,
		ServiceTime: time.Now().Unix(),
	}
}

package pack

import (
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

func BuildBaseResponse(statusCode int64, statusMsg string) *video.BaseResponse {
	return &video.BaseResponse{
		StatusCode:  statusCode,
		StatusMsg:   statusMsg,
		ServiceTime: time.Now().Unix(),
	}
}

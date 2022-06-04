package pack

import (
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
)

func BuildBaseResponse(statusCode int64, statusMsg string) *favorite.BaseResponse {
	return &favorite.BaseResponse{
		StatusCode:  statusCode,
		StatusMsg:   statusMsg,
		ServiceTime: time.Now().Unix(),
	}
}

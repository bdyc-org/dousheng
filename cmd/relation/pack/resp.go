package pack

import (
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/relation"
)

func BuildBaseResponse(statusCode int64, statusMsg string) *relation.BaseResponse {
	return &relation.BaseResponse{
		StatusCode:  statusCode,
		StatusMsg:   statusMsg,
		ServiceTime: time.Now().Unix(),
	}
}

package handlers

import (
	error2 "github.com/bdyc-org/dousheng/pkg/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int64       `json:"status_code"`
	Msg  string      `json:"status_msg"`
	Data interface{} `json:"data"`
}

type VideoParam struct {
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := error2.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code: Err.ErrCode,
		Msg:  Err.ErrMsg,
		Data: data,
	})
}

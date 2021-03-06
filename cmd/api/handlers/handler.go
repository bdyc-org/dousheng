package handlers

import (
	"net/http"

	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func SendErrResponse(c *gin.Context, statusCode int64, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  err.Error(),
	})
}

// relation
type RelaParam struct {
	ToUserId   int64  `json:"to_user_id" form:"to_user_id"`
	ActionType int64  `json:"action_type" form:"action_type"`
	Token      string `json:"token" form:"token"`
}
type RelaResponse struct {
	Code     int64       `json:"status_code"`
	Message  string      `json:"status_msg"`
	UserList interface{} `json:"user_list"`
}

func SendRelaResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	klog.Infof(Err.ErrMsg)
	c.JSON(http.StatusOK, RelaResponse{
		Code:     Err.ErrCode,
		Message:  Err.ErrMsg,
		UserList: data,
	})
}

type FollowParam struct {
	UserId int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

type Response struct {
	Code int32       `json:"status_code"`
	Msg  string      `json:"status_msg"`
	Data interface{} `json:"data"`
}

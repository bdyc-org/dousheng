package handlers

import (
	"net/http"

	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

func SendErrResponse(c *gin.Context, statusCode int64, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  err.Error(),
	})
}

type Response struct {
	Code    	int64       `json:"status_code"`
	Message 	string      `json:"status_msg"`
	UserList	interface{} `json:"user_list"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    	Err.ErrCode,
		Message: 	Err.ErrMsg,
		UserList:   data,
	})
}

type RelaParam struct {
	UserId int64 `json:"user_id"`
	ToUserId int64 `json:"to_user_id"`
	ActionType int64 `json:"action_type"`
}

type FollowParam struct {
	UserId int64 `json:"user_id"`
}
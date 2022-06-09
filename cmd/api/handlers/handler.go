package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendErrResponse(c *gin.Context, statusCode int64, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  err.Error(),
	})
}

type Response struct {
	Code        int64       `json:"status_code"`
	Message     string      `json:"status_msg"`
	CommentList interface{} `json:"comment_list"`
}

type CommentParam struct {
	UserID      int64  `json:"user_id"`
	Token       string `json:"token"`
	VideoID     int64  `json:"video_id"`
	ActionType  int32  `json:"action_type"`
	CommentText string `json:"comment_text"`
}

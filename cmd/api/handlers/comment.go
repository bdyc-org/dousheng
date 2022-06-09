package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Comment(c *gin.Context) {
	var commentVar CommentParam

	//获取参数
	//temp_string := c.Query("user_id")
	//temp_int64, err := strconv.ParseInt(temp_string, 10, 64)
	//if err != nil {
	//	commentVar.UserID = 0
	//}
	//commentVar.UserID = temp_int64

	commentVar.Token = c.Query("token")
	commentVar.CommentText = c.Query("comment_text")
	temp_string := c.Query("video_id")
	temp_int64, err := strconv.ParseInt(temp_string, 10, 64)
	if err != nil {
		commentVar.VideoID = 0
	}
	commentVar.VideoID = temp_int64
	temp_string = c.Query("action_type")
	temp_int64, err = strconv.ParseInt(temp_string, 10, 64)
	if err != nil {
		commentVar.ActionType = 0
	}
	commentVar.ActionType = int32(temp_int64)

	//检查参数是否合法
	if commentVar.VideoID == 0 || (commentVar.ActionType != 1 && commentVar.ActionType != 2) || commentVar.CommentText == "" {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	statusCode, err := rpc.CommentOperation(context.Background(), &comment.CommentRequest{
		UserId:      commentVar.UserID,
		VideoId:     commentVar.VideoID,
		ActionType:  commentVar.ActionType,
		CommentText: commentVar.CommentText,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	switch commentVar.ActionType {
	case 1:
		err = errors.New("评论成功")
	case 2:
		err = errors.New("删除评论成功")
	default:
		err = errors.New("未定义的操作")
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  err.Error(),
	})
}

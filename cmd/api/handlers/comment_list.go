package handlers

import (
	"context"
	"errors"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListParam struct {
	VideoID int64  `json:"video_id"`
	Token   string `json:"token"`
}

func CommentList(c *gin.Context) {
	var commentListVar CommentListParam

	//获取参数
	temp_string := c.Query("user_id")
	temp_int64, err := strconv.ParseInt(temp_string, 10, 64)
	if err != nil {
		commentListVar.VideoID = 0
	}
	commentListVar.VideoID = temp_int64
	commentListVar.Token = c.Query("token")

	commentList, statusCode, err := rpc.CommentList(context.Background(), &comment.QueryCommentRequest{
		VideoId: commentListVar.VideoID,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	if len(commentList) == 0 {
		err = errors.New("该视频暂未无评论")
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code":  statusCode,
		"status_msg":   err.Error(),
		"comment_list": commentList,
	})
}

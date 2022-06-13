package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type CommentListParam struct {
	Token   string `json:"token" form:"token"`
	VideoID int64  `json:"video_id" form:"video_id"`
}

func CommentList(c *gin.Context) {
	var commentListVar CommentListParam
	var user_id int64
	var statusCode int64
	var err error

	//获取参数
	if err := c.ShouldBindQuery(&commentListVar); err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//检查参数是否合法
	if commentListVar.VideoID == 0 {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//Token鉴权
	claims, err := ParserToken(commentListVar.Token)
	if err != nil {
		user_id = 0
	} else {
		username := claims.Username
		user_id, statusCode, err = rpc.Authentication(context.Background(), &user.AuthenticationRequest{
			Username: username,
		})
		if err != nil || user_id == 0 {
			SendErrResponse(c, statusCode, err)
			return
		}
	}

	commentList, statusCode, err := rpc.CommentList(context.Background(), &comment.CommentListRequest{
		UserId:  user_id,
		VideoId: commentListVar.VideoID,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	if len(commentList) == 0 {
		err = errors.New("该视频暂无评论")
	} else {
		err = errors.New("获取评论列表成功")
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code":  statusCode,
		"status_msg":   err.Error(),
		"comment_list": commentList,
	})
}

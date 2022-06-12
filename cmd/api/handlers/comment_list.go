package handlers

import (
	"context"
	"errors"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListParam struct {
	VideoID int64  `json:"video_id"`
	Token   string `json:"token"`
}

func CommentList(c *gin.Context) {
	var commentListVar CommentListParam

	//获取参数
	if err := c.ShouldBindQuery(&commentListVar); err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//检查参数是否合法
	if commentListVar.VideoID == 0 || len(commentListVar.Token) == 0 {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//Token鉴权
	claims, err := ParserToken(commentListVar.Token)
	if err != nil {
		SendErrResponse(c, errno.TokenInvalidErrCode, errno.ErrTokenInvalid)
		return
	}
	username := claims.Username
	user_id, statusCode, err := rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: username,
	})
	if err != nil || user_id == 0 {
		SendErrResponse(c, statusCode, err)
		return
	}

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

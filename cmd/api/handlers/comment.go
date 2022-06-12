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

func Comment(c *gin.Context) {
	var commentVar CommentParam
	//获取参数
	if err := c.ShouldBindQuery(&commentVar); err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}
	//检查参数是否合法
	if commentVar.VideoID == 0 || (commentVar.ActionType != 1 && commentVar.ActionType != 2) || commentVar.CommentText == "" {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	// 调用RPC
	//Token鉴权
	claims, err := ParserToken(commentVar.Token)
	if err != nil {
		SendRelaResponse(c, errno.NewErrNo(errno.TokenInvalidErrCode, errno.ErrTokenInvalid.Error()), nil)
		return
	}
	username := claims.Username
	user_id, statusCode, err := rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: username,
	})
	if err != nil || user_id == 0 {
		SendRelaResponse(c, errno.NewErrNo(statusCode, err.Error()), nil)
		return
	}

	statusCode, err = rpc.CommentOperation(context.Background(), &comment.CommentRequest{
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

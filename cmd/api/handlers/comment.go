package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type CommentParam struct {
	Token       string  `json:"token" form:"token"`
	VideoID     int64   `json:"video_id" form:"video_id"`
	ActionType  int64   `json:"action_type" form:"action_type"`
	CommentText *string `json:"comment_text" form:"comment_text"`
	CommentId   *int64  `json:"comment_id" form:"comment_id"`
}

func Comment(c *gin.Context) {
	var commentVar CommentParam
	//获取参数
	if err := c.ShouldBindQuery(&commentVar); err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//检查参数是否合法
	if commentVar.VideoID == 0 || (commentVar.ActionType != 1 && commentVar.ActionType != 2) {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}
	if (commentVar.ActionType == 1 && commentVar.CommentText == nil) || (commentVar.ActionType == 2 && commentVar.CommentId == nil) {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//Token鉴权
	claims, err := ParserToken(commentVar.Token)
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

	comment, statusCode, err := rpc.CommentOperation(context.Background(), &comment.CommentRequest{
		UserId:      user_id,
		VideoId:     commentVar.VideoID,
		ActionType:  commentVar.ActionType,
		CommentText: commentVar.CommentText,
		CommentId:   commentVar.CommentId,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	statusCode, err = rpc.VideoComment(context.Background(), &video.CommentOperationRequest{
		VideoId:    commentVar.VideoID,
		ActionType: commentVar.ActionType,
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
		err = errors.New("未定义操作")
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  err.Error(),
		"comment":     comment,
	})

}

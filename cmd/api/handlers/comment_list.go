package handlers

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

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

	resp, err := rpc.CommentList(context.Background(), &comment.QueryCommentRequest{
		UserId: user_id,
		VideoId: commentListVar.VideoID,
	})
	if err != nil {
		SendCommListResponse(c, err, resp.CommentList)
		return
	}

	if len(resp.CommentList) == 0 {
		SendCommListResponse(c, errno.Success.WithMessage("该视频暂未无评论"), resp.CommentList)
		return
	}

	SendCommListResponse(c, errno.Success.WithMessage("获取评论列表成功"), resp.CommentList)
}

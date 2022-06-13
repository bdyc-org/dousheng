package handlers

import (
	"context"
	"errors"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Follow(c *gin.Context) {
	var relaParam RelaParam
	if err := c.ShouldBindQuery(&relaParam); err != nil {
		SendRelaResponse(c, errno.ParamErr.WithMessage("参数获取失败"), nil)
		return
	}
	if relaParam.ToUserId == 0 || len(relaParam.Token) == 0 {
		SendRelaResponse(c, errno.ParamErr.WithMessage("参数不正确"), nil)
		return
	}

	// 调用RPC
	//Token鉴权
	claims, err := ParserToken(relaParam.Token)
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

	if user_id == relaParam.ToUserId {
		SendErrResponse(c, errno.ParamErrCode, errors.New("您不能关注自己"))
		return
	}

	// 关注或取关
	resp, err := rpc.RelaFollow(context.Background(), &relation.FollowRequest{
		UserId:     user_id,
		ToUserId:   relaParam.ToUserId,
		ActionType: relaParam.ActionType,
	})
	if err != nil {
		SendRelaResponse(c, err, resp)
		return
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		SendRelaResponse(c, errno.ParamErr, nil)
		return
	}

	// 调用userClient
	res, err := rpc.UserFollow(context.Background(), &user.FollowOperationRequest{
		FollowId:   relaParam.ToUserId,
		FollowerId: user_id,
		ActionType: relaParam.ActionType,
	})
	if err != nil {
		SendRelaResponse(c, err, resp)
		return
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		SendRelaResponse(c, errno.ParamErr, nil)
		return
	}

	SendRelaResponse(c, errno.Success, res)
}

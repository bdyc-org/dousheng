package handlers

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

func QueryFollow(c *gin.Context) {
	var followParam FollowParam
	if err := c.ShouldBindQuery(&followParam); err != nil {
		SendRelaResponse(c, errno.ParamErr.WithMessage("参数获取失败"), nil)
		return
	}
	if followParam.UserId == 0 || len(followParam.Token) == 0 {
		SendRelaResponse(c, errno.ParamErr.WithMessage("参数不正确"), nil)
		return
	}

	//Token鉴权
	claims, err := ParserToken(followParam.Token)
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

	var userList []*relation.User
	// 取ids
	FollowResp, err := rpc.QueryFollow(context.Background(), followParam.UserId)
	if err != nil {
		SendRelaResponse(c, errno.ServiceErr, nil)
		return
	}
	// 取userList
	resp, err := rpc.QueryUserList(context.Background(), &relation.QueryUserListRequest{
		UserId:  user_id,
		UserIds: FollowResp.FollowIds,
	})
	if err != nil {
		SendRelaResponse(c, errno.ServiceErr, nil)
		return
	}
	userList = resp.UserList
	SendRelaResponse(c, errno.Success.WithMessage("获取用户列表成功"), userList)
}

func QueryFollower(c *gin.Context) {
	var followParam FollowParam
	if err := c.ShouldBindQuery(&followParam); err != nil {
		SendRelaResponse(c, errno.ParamErr.WithMessage("参数获取失败"), nil)
		return
	}
	if followParam.UserId == 0 || len(followParam.Token) == 0 {
		SendRelaResponse(c, errno.ParamErr.WithMessage("参数不正确"), nil)
		return
	}

	//Token鉴权
	claims, err := ParserToken(followParam.Token)
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

	var userList []*relation.User
	// 取ids
	FollowerResp, err := rpc.QueryFollower(context.Background(), followParam.UserId)
	if err != nil {
		SendRelaResponse(c, errno.ServiceErr, nil)
		return
	}
	// 取userList
	resp, err := rpc.QueryUserList(context.Background(), &relation.QueryUserListRequest{
		UserId:  user_id,
		UserIds: FollowerResp.FollowerIds,
	})
	if err != nil {
		SendRelaResponse(c, errno.ServiceErr, nil)
		return
	}
	userList = resp.UserList
	SendRelaResponse(c, errno.Success.WithMessage("获取用户列表成功"), userList)
}

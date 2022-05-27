package handlers

import (
	"context"
	"strconv"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

func QueryFollow(c *gin.Context) {
	t1 := c.Query("user_id")
	ParmUserId, err := strconv.ParseInt(t1, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	println(ParmUserId)
	if ParmUserId == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	var userList []*relation.User
	// 取ids
	FollowResp , err := rpc.QueryFollow(context.Background(), ParmUserId)
	if err != nil {
		SendResponse(c, errno.ServiceErr, nil)
		return
	}
	// 取userList
	resp, err := rpc.QueryUserList(context.Background(), &relation.QueryUserListRequest{
		UserId: ParmUserId,
		UserIds: FollowResp.FollowIds,
	})
	if err != nil {
		SendResponse(c, errno.ServiceErr, nil)
		return
	}
	userList = resp.UserList
	SendResponse(c, errno.Success.WithMessage("获取用户列表成功"), userList)
}

func QueryFollower(c *gin.Context) {
	t1 := c.Query("user_id")
	ParmUserId, err := strconv.ParseInt(t1, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	if ParmUserId == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	var userList []*relation.User
	// 取ids
	FollowerResp , err := rpc.QueryFollower(context.Background(), ParmUserId)
	if err != nil {
		SendResponse(c, errno.ServiceErr, nil)
		return
	}
	// 取userList
	resp, err := rpc.QueryUserList(context.Background(), &relation.QueryUserListRequest{
		UserId: ParmUserId,
		UserIds: FollowerResp.FollowerIds,
	})
	if err != nil {
		SendResponse(c, errno.ServiceErr, nil)
		return
	}
	userList = resp.UserList
	SendResponse(c, errno.Success.WithMessage("获取用户列表成功"), userList)
}
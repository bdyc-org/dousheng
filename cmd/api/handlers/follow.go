package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

func QueryFollow(c *gin.Context) {
	t1 := c.Query("user_id")
	ParmUserId, err := strconv.ParseInt(t1, 10, 64)
	if err != nil {
		ParmUserId = 0
	}
	token := c.Query("token")
	if ParmUserId == 0 || len(token) == 0 {
		SendErrResponse(c, errno.ParamErrCode, errors.New(errno.Errparameter.Error()))
		return
	}

	//Token鉴权
	claims, err := ParserToken(token)
	if err != nil {
		SendErrResponse(c, errno.TokenInvalidErrCode, errno.ErrTokenInvalid)
		return
	}
	username := claims.Username
	user_id, statusCode, err := rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: username,
	})
	if err != nil || user_id == 0 || user_id != ParmUserId {
		SendErrResponse(c, statusCode, err)
		return
	}

	var userList []*relation.User
	// 取ids
	FollowResp , err := rpc.QueryFollow(context.Background(), ParmUserId)
	if err != nil {
		SendErrResponse(c, errno.ServiceErrCode, err)
		return
	}
	// 取userList
	resp, err := rpc.QueryUserList(context.Background(), &relation.QueryUserListRequest{
		UserId: ParmUserId,
		UserIds: FollowResp.FollowIds,
	})
	if err != nil {
		SendErrResponse(c, errno.ServiceErrCode, err)
		return
	}
	userList = resp.UserList
	
	c.JSON(http.StatusOK, gin.H{
		"status_code": 	errno.SuccessCode,
		"status_msg":  	"获取用户列表成功",
		"user_list":    userList,
	})
}

func QueryFollower(c *gin.Context) {
	t1 := c.Query("user_id")
	ParmUserId, err := strconv.ParseInt(t1, 10, 64)
	if err != nil {
		ParmUserId = 0
	}
	token := c.Query("token")
	if ParmUserId == 0 || len(token) == 0 {
		SendErrResponse(c, errno.ParamErrCode, errors.New(errno.Errparameter.Error()))
		return
	}

	//Token鉴权
	claims, err := ParserToken(token)
	if err != nil {
		SendErrResponse(c, errno.TokenInvalidErrCode, errno.ErrTokenInvalid)
		return
	}
	username := claims.Username
	user_id, statusCode, err := rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: username,
	})
	if err != nil || user_id == 0 || user_id != ParmUserId {
		SendErrResponse(c, statusCode, err)
		return
	}

	var userList []*relation.User
	FollowResp , err := rpc.QueryFollower(context.Background(), ParmUserId)
	if err != nil {
		SendErrResponse(c, errno.ServiceErrCode, err)
		return
	}
	// 取userList
	resp, err := rpc.QueryUserList(context.Background(), &relation.QueryUserListRequest{
		UserId: ParmUserId,
		UserIds: FollowResp.FollowerIds,
	})
	if err != nil {
		SendErrResponse(c, errno.ServiceErrCode, err)
		return
	}
	userList = resp.UserList
	
	c.JSON(http.StatusOK, gin.H{
		"status_code": 	errno.SuccessCode,
		"status_msg":  	"获取用户列表成功",
		"user_list":    userList,
	})
}
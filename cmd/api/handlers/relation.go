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

func Follow(c *gin.Context) {
	t1 := c.Query("user_id")
	ParmUserId, err := strconv.ParseInt(t1, 10, 64)
	if err != nil {
		ParmUserId = 0
	}
	t2 := c.Query("to_user_id")
	ParmToUserId, err := strconv.ParseInt(t2, 10, 64)
	if err != nil {
		ParmToUserId = 0
	}
	t3 := c.Query("action_type")
	ParmActionType, err := strconv.ParseInt(t3, 10, 64)
	if err != nil {
		ParmActionType = 0
	}
	token := c.Query("token")
	if ParmUserId == 0 || ParmToUserId == 0 || ParmActionType == 0 || len(token) == 0 {
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

	// 调用RPC
	resp, err := rpc.RelaFollow(context.Background(), &relation.FollowRequest{
		UserId: ParmUserId,
		ToUserId: ParmToUserId,
		ActionType: ParmActionType,
	})
	if err != nil {
		SendErrResponse(c, resp.BaseResp.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code":	resp.BaseResp.StatusCode,
		"status_msg":	resp.BaseResp.StatusMsg,
	})
}
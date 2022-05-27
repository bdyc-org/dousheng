package handlers

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Follow(c *gin.Context) {
	var relaParam RelaParam
	if err := c.ShouldBind(&relaParam); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	if relaParam.UserId == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	// 调用RPC
	resp, err := rpc.RelaFollow(context.Background(), &relation.FollowRequest{
		UserId: relaParam.UserId,
		ToUserId: relaParam.ToUserId,
		ActionType: relaParam.ActionType,
	})
	if err != nil {
		SendResponse(c, err, resp)
		return
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	SendResponse(c, errno.Success, resp)
}
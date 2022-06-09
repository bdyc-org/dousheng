package handlers

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	error2 "github.com/bdyc-org/dousheng/pkg/error"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PublishList(c *gin.Context) {
	user_id, err := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}

	token := c.PostForm("token")
	//TODO token judge

	req := &video.DouyinPublishListRequest{
		UserId: user_id,
		Token:  token,
	}

	_, err = rpc.PublishList(context.Background(), req)
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}

	SendResponse(c, error2.Success, map[string]interface{}{})

}

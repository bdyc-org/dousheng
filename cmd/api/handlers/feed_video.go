package handlers

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	error2 "github.com/bdyc-org/dousheng/pkg/error"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FeedVideo(c *gin.Context) {
	latest_time, err := strconv.ParseInt(c.PostForm("latest_time"), 10, 64)

	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}

	token := c.PostForm("token")
	//TODO token judge
	req := &video.DouyinFeedRequest{
		LatestTime: &latest_time,
		Token:      &token,
	}
	_, _, err = rpc.FeedVideo(context.Background(), req)

	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}

	SendResponse(c, error2.Success, map[string]interface{}{})

}

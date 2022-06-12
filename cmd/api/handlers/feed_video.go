package handlers

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	error2 "github.com/bdyc-org/dousheng/pkg/error"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FeedVideo(c *gin.Context) {
	latest_time, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)

	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}

	token := c.Query("token")
	claims, err := ParserToken(token)
	username := claims.Username
	user_id, statusCode, err := rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: username,
	})
	if err != nil || user_id == 0 {
		SendErrResponse(c, statusCode, err)
		return
	}

	req := &video.DouyinFeedRequest{
		LatestTime: &latest_time,
		UserId:     &user_id,
	}
	_, _, err = rpc.FeedVideo(context.Background(), req)

	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}

	SendResponse(c, error2.Success, map[string]interface{}{})

}

package handlers

import (
	"context"
	"fmt"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/cmd/video/pack"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	error2 "github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	VideoResponse
	VideoList []*video.Video `json:"video_list"`
}

func PublishList(c *gin.Context) {
	user_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}

	token := c.Query("token")
	fmt.Println(token)
	claims, err := ParserToken(token)
	username := claims.Username
	user_id, statusCode, err := rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: username,
	})
	if err != nil || user_id == 0 {
		SendErrResponse(c, statusCode, err)
		return
	}

	ipvf, err := pack.GetLocalIPv4Address()
	if err != nil {
		panic(err)
	}
	videos, err := rpc.PublishList(context.Background(), &video.DouyinPublishListRequest{
		UserId: user_id,
	})
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}

	for _, video_ := range videos {
		video_.PlayUrl = "http://" + ipvf + ":8080/static/videos/" + video_.PlayUrl
	}
	c.JSON(http.StatusOK, VideoListResponse{
		VideoResponse: VideoResponse{
			StatusCode: 0,
		},
		VideoList: videos,
	})

}

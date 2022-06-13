package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type FeedParam struct {
	LatestTime int64  `json:"latest_time" form:"latest_time"`
	Token      string `json:"token" form:"token"`
}

func Feed(c *gin.Context) {
	// var feedVar FeedParam
	// var user_id int64
	// //获取参数
	// if err := c.ShouldBindQuery(&feedVar); err != nil {
	// 	SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
	// 	return
	// }

	// fmt.Println(feedVar.Token)
	// fmt.Println(feedVar.LatestTime)
	// //检查参数
	// if feedVar.LatestTime == 0 {
	// 	feedVar.LatestTime = time.Now().Unix()
	// }
	// if len(feedVar.Token) == 0 {
	// 	user_id = 0
	// } else {
	// 	//Token鉴权
	// 	var statusCode int64
	// 	claims, err := ParserToken(feedVar.Token)
	// 	if err != nil {
	// 		SendErrResponse(c, errno.TokenInvalidErrCode, errno.ErrTokenInvalid)
	// 		return
	// 	}
	// 	username := claims.Username
	// 	user_id, statusCode, err = rpc.Authentication(context.Background(), &user.AuthenticationRequest{
	// 		Username: username,
	// 	})
	// 	if err != nil {
	// 		SendErrResponse(c, statusCode, err)
	// 		return
	// 	}
	// }
	_, videoList, _, _ := rpc.Feed(context.Background(), &video.FeedRequest{
		LatestTime: time.Now().Unix(),
		UserId:     0,
	})

	// nextTime, videoList, statusCode, err := rpc.Feed(context.Background(), &video.FeedRequest{
	// 	LatestTime: time.Now().Unix(),
	// 	UserId:     0,
	// })
	// if err != nil {
	// 	SendErrResponse(c, statusCode, err)
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "获取feed流成功",
		// "next_time":   nextTime,
		"video_list": videoList,
	})
}

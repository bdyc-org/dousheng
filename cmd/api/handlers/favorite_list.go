package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type FavoriteListParam struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func FacoriteList(c *gin.Context) {
	var favoriteListVar FavoriteListParam

	//获取参数
	temp_string := c.Query("user_id")
	temp_int64, err := strconv.ParseInt(temp_string, 10, 64)
	if err != nil {
		favoriteListVar.UserID = 0
	}
	favoriteListVar.UserID = temp_int64
	favoriteListVar.Token = c.Query("token")

	//Token鉴权
	claims, err := ParserToken(favoriteListVar.Token)
	if err != nil {
		SendErrResponse(c, errno.TokenInvalidErrCode, errno.ErrTokenInvalid)
		return
	}
	username := claims.Username
	user_id, statusCode, err := rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: username,
	})
	if err != nil || user_id == 0 {
		SendErrResponse(c, statusCode, err)
		return
	}

	videoList, statusCode, err := rpc.FacoriteList(context.Background(), &favorite.FavoriteListRequest{
		UserId: user_id,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	if len(videoList) == 0 {
		err = errors.New("您还没有喜欢的视频")
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  err.Error(),
		"video_list":  videoList,
	})
}

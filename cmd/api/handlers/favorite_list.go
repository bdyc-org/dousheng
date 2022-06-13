package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type FavoriteListParam struct {
	UserID int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

func FavoriteList(c *gin.Context) {
	var favoriteListVar FavoriteListParam

	//获取参数
	if err := c.ShouldBindQuery(&favoriteListVar); err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//检查参数是否合法
	if favoriteListVar.UserID == 0 || len(favoriteListVar.Token) == 0 {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

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

	videoList, statusCode, err := rpc.FavoriteList(context.Background(), &favorite.FavoriteListRequest{
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

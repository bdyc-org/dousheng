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

type FavoriteParam struct {
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
	VideoID    int64  `json:"video_id"`
	ActionType int64  `json:"action_type"`
}

func Favorite(c *gin.Context) {
	var favoriteVar FavoriteParam

	//获取参数
	// temp_string := c.Query("user_id")
	// temp_int64, err := strconv.ParseInt(temp_string, 10, 64)
	// if err != nil {
	// 	favoriteVar.UserID = 0
	// }
	// favoriteVar.UserID = temp_int64
	favoriteVar.Token = c.Query("token")
	temp_string := c.Query("video_id")
	temp_int64, err := strconv.ParseInt(temp_string, 10, 64)
	if err != nil {
		favoriteVar.VideoID = 0
	}
	favoriteVar.VideoID = temp_int64
	temp_string = c.Query("action_type")
	temp_int64, err = strconv.ParseInt(temp_string, 10, 64)
	if err != nil {
		favoriteVar.ActionType = 0
	}
	favoriteVar.ActionType = temp_int64

	//检查参数是否合法
	if favoriteVar.VideoID == 0 || (favoriteVar.ActionType != 1 && favoriteVar.ActionType != 2) {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//Token鉴权
	claims, err := ParserToken(favoriteVar.Token)
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

	statusCode, err = rpc.FavoriteOperation(context.Background(), &favorite.FavoriteOperationRequest{
		UserId:     user_id,
		VideoId:    favoriteVar.VideoID,
		ActionType: favoriteVar.ActionType,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	switch favoriteVar.ActionType {
	case 1:
		err = errors.New("点赞成功，感谢您的支持~")
	case 2:
		err = errors.New("取消点赞成功")
	default:
		err = errors.New("未定义的操作")
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  err.Error(),
	})
}

package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type FavoriteParam struct {
	Token      string `json:"token" form:"token"`
	VideoID    int64  `json:"video_id" form:"video_id"`
	ActionType int64  `json:"action_type" form:"action_type"`
}

func Favorite(c *gin.Context) {
	var favoriteVar FavoriteParam

	//获取参数
	if err := c.ShouldBindQuery(&favoriteVar); err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

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

	videos, statusCode, err := rpc.MGetVideo(context.Background(), &video.MGetVideoRequest{
		UserId:   user_id,
		VideoIds: []int64{favoriteVar.VideoID},
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	statusCode, err = rpc.UserFavorite(context.Background(), &user.FavoriteOperationRequest{
		UserId:      user_id,
		VideoAuthor: videos[0].Author.Id,
		ActionType:  favoriteVar.ActionType,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	statusCode, err = rpc.VideoFavorite(context.Background(), &video.FavoriteOperationRequest{
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

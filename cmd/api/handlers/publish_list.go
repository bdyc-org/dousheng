package handlers

import (
	"context"
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type PublishListParam struct {
	Token  string `json:"token" form:"token"`
	UserID int64  `json:"user_id" form:"user_id"`
}

func PublishList(c *gin.Context) {
	var publishListVar PublishListParam
	if err := c.ShouldBindQuery(&publishListVar); err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//检查参数是否合法
	if len(publishListVar.Token) == 0 || publishListVar.UserID == 0 {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//Token鉴权
	claims, err := ParserToken(publishListVar.Token)
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

	videoList, statusCode, err := rpc.PublishList(context.Background(), &video.PublishListRequest{
		UserId:   user_id,
		AuthorId: publishListVar.UserID,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "获取发布列表成功",
		"video_list":  videoList,
	})

}

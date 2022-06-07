package handlers

import (
	"context"
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type UserInfoParam struct {
	UserID int64  `json:"user_id" form:"user_id"`
	Token  string `json:"token" form:"token"`
}

func UserInfo(c *gin.Context) {
	var userInfoVar UserInfoParam

	//获取参数
	if err := c.ShouldBindQuery(&userInfoVar); err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//检查参数是否合法
	if userInfoVar.UserID == 0 || len(userInfoVar.Token) == 0 {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//Token鉴权
	claims, err := ParserToken(userInfoVar.Token)
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

	//获取用户信息
	users, statusCode, err := rpc.MGetUser(context.Background(), &user.MGetUserRequest{
		UserId:  user_id,
		UserIds: []int64{userInfoVar.UserID},
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	if len(users) == 0 {
		SendErrResponse(c, errno.ServiceErrCode, errno.ErrService)
		return
	}
	user := users[0]

	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "获取用户信息成功",
		"user":        user,
	})

}

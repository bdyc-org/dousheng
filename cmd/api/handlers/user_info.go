package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type UserInfoParm struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func UserInfo(c *gin.Context) {
	var userInfoVar UserInfoParm

	temp := c.Query("user_id")
	temp_id, err := strconv.ParseInt(temp, 10, 64)
	if err != nil {
		userInfoVar.UserID = 0
	}
	userInfoVar.UserID = temp_id
	userInfoVar.Token = c.Query("token")

	fmt.Println(userInfoVar)

	if userInfoVar.UserID == 0 || len(userInfoVar.Token) == 0 {
		// SendLoginResponse(c, errno.ParamErr)
		return
	}

	claims, err := ParserToken(userInfoVar.Token)
	if err != nil {
		return
	}
	username := claims.Username
	user_id, err := rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: username,
	})
	if err != nil {
		return
	}

	users, err := rpc.MGetUser(context.Background(), &user.MGetUserRequest{
		UserId:  user_id,
		UserIds: []int64{userInfoVar.UserID},
	})
	if err != nil {
		return
	}

	if len(users) == 0 {
		return
	}
	user := users[0]

	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "获取用户信息成功",
		"user":        user,
	})

}

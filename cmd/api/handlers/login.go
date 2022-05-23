package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type LoginParm struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func SendLoginResponse(c *gin.Context, err errno.ErrNo) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": err.ErrCode,
		"status_msg":  err.ErrMsg,
		"user_id":     0,
		"token":       "",
	})
}

func Login(c *gin.Context) {
	var loginVar LoginParm

	loginVar.UserName = c.Query("username")
	loginVar.PassWord = c.Query("password")

	fmt.Println(loginVar)

	if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
		SendLoginResponse(c, errno.ParamErr)
		return
	}

	user_id, err := rpc.CheckUser(context.Background(), &user.CheckUserRequest{
		Username: loginVar.UserName,
		Password: loginVar.PassWord,
	})
	if err != nil {
		SendLoginResponse(c, errno.ConvertErr(err))
		return
	}
	token, err := GenerateToken(loginVar.UserName)
	if err != nil {
		SendLoginResponse(c, errno.ConvertErr(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "登录成功",
		"user_id":     user_id,
		"token":       token,
	})

}

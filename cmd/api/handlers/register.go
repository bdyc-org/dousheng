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

type RegisterParm struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func SendRegisterResponse(c *gin.Context, err errno.ErrNo) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": err.ErrCode,
		"status_msg":  err.ErrMsg,
		"user_id":     0,
		"token":       "",
	})
}

func Register(c *gin.Context) {
	var registerVar RegisterParm

	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	fmt.Println(registerVar)

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendRegisterResponse(c, errno.ParamErr)
		return
	}

	user_id, err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		fmt.Println("我来了4")
		SendRegisterResponse(c, errno.ConvertErr(err))
		return
	}
	token, err := GenerateToken(registerVar.UserName)
	if err != nil {
		SendRegisterResponse(c, errno.ConvertErr(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "注册成功",
		"user_id":     user_id,
		"token":       token,
	})

}

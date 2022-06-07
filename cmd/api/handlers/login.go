package handlers

import (
	"context"
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type LoginParam struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

func Login(c *gin.Context) {
	var loginVar LoginParam
	//获取参数
	loginVar.UserName = c.Query("username")
	loginVar.PassWord = c.Query("password")

	//检查参数是否合法
	if len(loginVar.UserName) == 0 || len(loginVar.PassWord) < 6 {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//数据库查询用户名，密码，成功返回用户id
	user_id, statusCode, err := rpc.CheckUser(context.Background(), &user.CheckUserRequest{
		Username: loginVar.UserName,
		Password: loginVar.PassWord,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	//生成Token
	token, err := GenerateToken(loginVar.UserName)
	if err != nil {
		SendErrResponse(c, errno.ServiceErrCode, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "登录成功",
		"user_id":     user_id,
		"token":       token,
	})

}

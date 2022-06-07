package handlers

import (
	"context"
	"net/http"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/gin-gonic/gin"
)

type RegisterParam struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

func Register(c *gin.Context) {
	var registerVar RegisterParam

	//获取参数
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	//检查参数是否合法
	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) < 6 {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//将注册信息写入数据库
	user_id, statusCode, err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
		return
	}

	//生成Token
	token, err := GenerateToken(registerVar.UserName)
	if err != nil {
		SendErrResponse(c, errno.ServiceErrCode, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "注册成功",
		"user_id":     user_id,
		"token":       token,
	})

}

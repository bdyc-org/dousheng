package handlers

import (
	"context"
	"fmt"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	error2 "github.com/bdyc-org/dousheng/pkg/error"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PublishList(c *gin.Context) {
	user_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}

	token := c.Query("token")
	fmt.Println(token)
	//claims, err := ParserToken(token)
	//username := claims.Username
	//user_id, statusCode, err := rpc.Authentication(context.Background(), &user.AuthenticationRequest{
	//	Username: username,
	//})
	//if err != nil || user_id == 0 {
	//	SendErrResponse(c, statusCode, err)
	//	return
	//}

	_, err = rpc.PublishList(context.Background(), &video.DouyinPublishListRequest{
		UserId: user_id,
	})
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}

	SendResponse(c, error2.Success, map[string]interface{}{})

}

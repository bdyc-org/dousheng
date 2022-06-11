package handlers

import (
	"context"
	"fmt"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/constants"
	error2 "github.com/bdyc-org/dousheng/pkg/error"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func PublishVideo(c *gin.Context) {
	token := c.PostForm("token")
	claims, err := ParserToken(token)
	username := claims.Username
	user_id, statusCode, err := rpc.Authentication(context.Background(), &user.AuthenticationRequest{
		Username: username,
	})
	if err != nil || user_id == 0 {
		SendErrResponse(c, statusCode, err)
		return
	}

	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
		return
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", claims.Username, filename)
	saveFile := filepath.Join("../../../public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
		return
	}

	req := &video.DouyinPublishActionRequest{
		FileName: filename,
		UserId:   user_id,
		Title:    title,
	}
	err = rpc.PublicVideo(context.Background(), req)
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}
	SendResponse(c, error2.Success, map[string]interface{}{constants.Title: title, constants.Videos: finalName})
}

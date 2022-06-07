package handlers

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/constants"
	error2 "github.com/bdyc-org/dousheng/pkg/error"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func PublishVideo(c *gin.Context) {
	//TODO token judge
	token := c.PostForm("token")
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
		return
	}
	filename := filepath.Base(data.Filename)
	//TODO  filaname need to be changed
	//finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("../../../public/", filename)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
		return
	}
	req := &video.DouyinPublishActionRequest{
		FileName: filename,
		Token:    token,
		Title:    title,
	}
	err = rpc.PublicVideo(context.Background(), req)
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
	}
	SendResponse(c, error2.Success, map[string]interface{}{constants.Title: title, constants.Videos: filename})
}

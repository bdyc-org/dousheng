package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bdyc-org/dousheng/cmd/api/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func PublishVideo(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	//检查参数是否合法
	if len(token) == 0 || len(title) == 0 {
		SendErrResponse(c, errno.ParamErrCode, errno.Errparameter)
		return
	}

	// token鉴权
	claims, err := ParserToken(token)
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

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user_id, filename)
	saveFile := filepath.Join("./public/videos/", finalName)
	fmt.Println(saveFile)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		SendErrResponse(c, errno.ServiceErrCode, errno.ErrService)
		return
	}

	fileNameSplit := strings.Split(finalName, ".")
	coverName := fileNameSplit[0]

	reader := ExampleReadFrameAsJpeg("./public/videos/"+finalName, 5)
	img, err := imaging.Decode(reader)
	if err != nil {
		SendErrResponse(c, errno.ServiceErrCode, errno.ErrService)
	}
	err = imaging.Save(img, "./public/cover/"+coverName+".jpeg")
	if err != nil {
		SendErrResponse(c, errno.ServiceErrCode, errno.ErrService)
	}

	statusCode, err = rpc.CreateVideo(context.Background(), &video.CreateVideoRequest{
		AuthorId: user_id,
		PlayUrl:  "videos/" + finalName,
		CoverUrl: "cover/" + coverName + ".jpeg",
		Title:    title,
	})
	if err != nil {
		SendErrResponse(c, statusCode, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": errno.SuccessCode,
		"status_msg":  "投稿成功",
	})
}

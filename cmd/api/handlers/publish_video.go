package handlers

import (
	"fmt"
	error2 "github.com/bdyc-org/dousheng/pkg/error"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func PublishVideo(c *gin.Context) {
	//TODO token judge
	token := c.PostForm("token")
	title := c.PostForm("title")
	fmt.Println(token)
	data, err := c.FormFile("data")
	if err != nil {
		SendResponse(c, error2.ConvertErr(err), nil)
		return
	}
	filename := filepath.Base(data.Filename)
	//TODO  filaname need to be changed
	//finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("../../../public/", filename)

}

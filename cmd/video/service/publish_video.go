package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"io/ioutil"
	"os/exec"
)

type PublishVideoService struct {
	ctx context.Context
}

// NewPublishVideoService new PublishVideoService
func NewPublishVideoService(ctx context.Context) *PublishVideoService {
	return &PublishVideoService{ctx: ctx}
}

func (v *PublishVideoService) PublishVideo(req *video.DouyinPublishActionRequest) error {
	covername, err := GetFrame(req.FileName)

	if err != nil {
		panic(err)
	}
	fmt.Println(req.FileName)
	video := &db.Video{
		Title:          req.Title,
		Play_url:       req.FileName,
		Cover_url:      covername,
		User_id:        uint(req.UserId),
		Favorite_count: 0,
		Comment_count:  0,
	}

	return db.PublishVideo(v.ctx, []*db.Video{video})
}

func GetFrame(filename string) (covername string, err error) {
	width := 2752
	height := 2208
	//获取任意帧数
	// cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", strconv.Itoa(index), "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")

	buf := new(bytes.Buffer)

	cmd.Stdout = buf

	if cmd.Run() != nil {
		panic("could not generate frame")
	}
	//cover name == c + filename
	err = ioutil.WriteFile("./public/cover/"+"c"+filename, buf.Bytes(), 0666)
	covername = "c" + filename
	return covername, err
}

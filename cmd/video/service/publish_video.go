package service

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

type PublishVideoService struct {
	ctx context.Context
}

// NewPublishVideoService new PublishVideoService
func NewPublishVideoService(ctx context.Context) *PublishVideoService {
	return &PublishVideoService{ctx: ctx}
}

func (v *PublishVideoService) PublishVideo(req *video.DouyinPublishActionRequest) error {
	//TODO to obtain user_id

	video := &db.Video{
		Title:          req.Title,
		Play_url:       req.FileName,
		Cover_url:      req.FileName + "1",
		User_id:        0,
		Favorite_count: 0,
		Comment_count:  0,
	}

	return db.PublishVideo(v.ctx, []*db.Video{video})
}
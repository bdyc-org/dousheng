package service

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/cmd/video/pack"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

type PublishListService struct {
	ctx context.Context
}

func NewPublishListService(ctx context.Context) *PublishVideoService {
	return &PublishVideoService{ctx: ctx}
}

func (v *PublishVideoService) PublishList(req *video.DouyinPublishListRequest) ([]*video.Video, error) {
	videos, err := db.QueryVideos(v.ctx, uint(req.UserId))

	if err != nil {
		return nil, err
	}

	return pack.Videos(videos), nil
}

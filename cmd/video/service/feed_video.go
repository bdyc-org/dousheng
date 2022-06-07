package service

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/video/dal/db"

	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

type FeedVideoService struct {
	ctx context.Context
}

func NewFeedVideoService(ctx context.Context) *FeedVideoService {
	return &FeedVideoService{ctx: ctx}
}

func (v *FeedVideoService) FeedVideoService(req *video.DouyinFeedRequest) ([]*db.Video, *int64, error) {
	return db.VideoFeed(v.ctx, req.LatestTime)
}

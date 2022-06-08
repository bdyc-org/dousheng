package service

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/cmd/video/pack"

	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

type FeedVideoService struct {
	ctx context.Context
}

func NewFeedVideoService(ctx context.Context) *FeedVideoService {
	return &FeedVideoService{ctx: ctx}
}

func (v *FeedVideoService) FeedVideo(req *video.DouyinFeedRequest) ([]*video.Video, *int64, error) {
	videos, nextTime, err := db.VideoFeed(v.ctx, req.LatestTime)

	if err != nil {
		return nil, nil, err
	}
	return pack.Videos(videos), nextTime, nil
}

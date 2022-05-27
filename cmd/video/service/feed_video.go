package service

import (
	"context"

	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

type FeedVideoService struct {
	ctx context.Context
}

func NewFeedVideoService(ctx context.Context) *FeedVideoService {
	return &FeedVideoService{ctx: ctx}
}

func (v *FeedVideoService) FeedVideoService(req *video.DouyinFeedRequest) error {
	return nil
}

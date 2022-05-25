package main

import (
	"context"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// FeedVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FeedVideo(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

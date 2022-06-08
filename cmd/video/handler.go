package main

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/video/service"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	error2 "github.com/bdyc-org/dousheng/pkg/error"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// FeedVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FeedVideo(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	resp = new(video.DouyinFeedResponse)
	if req.LatestTime == nil {
		resp.StatusCode = error2.ParamErr.ErrCode
		resp.StatusMsg = &error2.ParamErr.ErrMsg
	}

	resp.VideoList, resp.NextTime, err = service.NewFeedVideoService(ctx).FeedVideo(req)
	return
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	resp = new(video.DouyinPublishActionResponse)
	if len(req.Title) == 0 || len(req.FileName) == 0 {
		resp.StatusCode = error2.ParamErr.ErrCode
		resp.StatusMsg = &error2.ParamErr.ErrMsg
	}
	err = service.NewPublishVideoService(ctx).PublishVideo(req)
	if err != nil {
		resp.StatusCode = error2.ConvertErr(err).ErrCode
		return resp, nil
	}
	resp.StatusCode = error2.Success.ErrCode
	resp.StatusMsg = &error2.Success.ErrMsg
	return resp, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

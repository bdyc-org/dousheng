package main

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/video/service"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	error2 "github.com/bdyc-org/dousheng/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// FeedVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FeedVideo(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	resp = new(video.DouyinFeedResponse)
	if req.LatestTime == nil {
		resp.StatusCode = int32(error2.ParamErr.ErrCode)
		resp.StatusMsg = &error2.ParamErr.ErrMsg
		return resp, err
	}

	resp.VideoList, resp.NextTime, err = service.NewFeedVideoService(ctx).FeedVideo(req)
	resp.StatusCode = int32(error2.Success.ErrCode)
	resp.StatusMsg = &error2.Success.ErrMsg
	return resp, nil
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	resp = new(video.DouyinPublishActionResponse)
	if len(req.Title) == 0 || len(req.FileName) == 0 {
		resp.StatusCode = int32(error2.ParamErr.ErrCode)
		resp.StatusMsg = &error2.ParamErr.ErrMsg
		return resp, err
	}
	err = service.NewPublishVideoService(ctx).PublishVideo(req)
	if err != nil {
		resp.StatusCode = int32(error2.ConvertErr(err).ErrCode)
		return resp, err
	}
	resp.StatusCode = int32(error2.Success.ErrCode)
	resp.StatusMsg = &error2.Success.ErrMsg
	return resp, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	resp = new(video.DouyinPublishListResponse)
	if req.UserId == 0 {
		resp.StatusCode = int32(error2.ParamErr.ErrCode)
		resp.StatusMsg = &error2.ParamErr.ErrMsg
		return resp, err
	}

	videos, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp.StatusCode = int32(error2.ConvertErr(err).ErrCode)
		return resp, err
	}
	resp.StatusCode = int32(error2.Success.ErrCode)
	resp.StatusMsg = &error2.Success.ErrMsg
	resp.VideoList = videos
	return resp, nil
}

// VideoFavorite implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoFavorite(ctx context.Context, req *video.DouyinVideoFavoriteRequest) (resp *video.DouyinVideoFavoriteResponse, err error) {
	resp = new(video.DouyinVideoFavoriteResponse)
	if req.VideoId == 0 || (req.Action != 1 && req.Action != 2) {
		resp.StatusCode = int32(error2.ParamErr.ErrCode)
		resp.StatusMsg = &error2.ParamErr.ErrMsg
		return resp, err
	}

	err = service.NewVideoFavoriteService(ctx).VideoFavorite(req)

	if err != nil {
		resp.StatusCode = int32(error2.ConvertErr(err).ErrCode)
		return resp, err
	}
	resp.StatusCode = int32(error2.Success.ErrCode)
	resp.StatusMsg = &error2.Success.ErrMsg
	return resp, nil
}

// VideoComment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoComment(ctx context.Context, req *video.DouyinVideoCommentRequest) (resp *video.DouyinVideoCommentResponse, err error) {
	resp = new(video.DouyinVideoCommentResponse)
	if req.VideoId == 0 || (req.Action != 1 && req.Action != 2) {
		resp.StatusCode = int32(error2.ParamErr.ErrCode)
		resp.StatusMsg = &error2.ParamErr.ErrMsg
		return resp, err
	}

	err = service.NewVideoCommentService(ctx).VideoComment(req)

	if err != nil {
		resp.StatusCode = int32(error2.ConvertErr(err).ErrCode)
		return resp, err
	}
	resp.StatusCode = int32(error2.Success.ErrCode)
	resp.StatusMsg = &error2.Success.ErrMsg
	return resp, nil
}

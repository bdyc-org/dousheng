package main

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/video/pack"
	"github.com/bdyc-org/dousheng/cmd/video/service"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	resp = new(video.FeedResponse)
	videos, nextTime, statusCode, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.NextTime = req.LatestTime
		resp.VideoList = nil
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取视频Feed流成功")
	resp.NextTime = nextTime
	resp.VideoList = videos
	return resp, nil
}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *video.CreateVideoRequest) (resp *video.CreateVideoResponse, err error) {
	// TODO: Your code here...
	resp = new(video.CreateVideoResponse)
	if req.AuthorId == 0 || len(req.CoverUrl) == 0 || len(req.PlayUrl) == 0 || len(req.Title) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		return resp, nil
	}
	statusCode, err := service.NewCreateVideoService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "投稿视频成功")
	return resp, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.PublishListResponse)
	if req.UserId == 0 || req.AuthorId == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		resp.VideoList = nil
		return resp, nil
	}
	videos, statusCode, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.VideoList = nil
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取投稿列表成功")
	resp.VideoList = videos
	return resp, nil
}

// MGetVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideo(ctx context.Context, req *video.MGetVideoRequest) (resp *video.MGetVideoResponse, err error) {
	// TODO: Your code here...
	resp = new(video.MGetVideoResponse)
	if len(req.VideoIds) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "该列表还没有视频")
		resp.VideoList = nil
		return resp, nil
	}
	videos, statusCode, err := service.NewMGetVideoService(ctx).MGetVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.VideoList = nil
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取视频列表成功")
	resp.VideoList = videos
	return resp, nil
}

// Favorite implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Favorite(ctx context.Context, req *video.FavoriteOperationRequest) (resp *video.FavoriteOperationResponse, err error) {
	// TODO: Your code here...
	resp = new(video.FavoriteOperationResponse)
	if req.VideoId == 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		return resp, nil
	}
	statusCode, err := service.NewFavoriteService(ctx).Favorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "赞操作成功")
	return resp, nil
}

// Comment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Comment(ctx context.Context, req *video.CommentOperationRequest) (resp *video.CommentOperationResponse, err error) {
	// TODO: Your code here...
	resp = new(video.CommentOperationResponse)
	if req.VideoId == 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		return resp, nil
	}
	statusCode, err := service.NewCommentService(ctx).Comment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "评论操作成功")
	return resp, nil
}

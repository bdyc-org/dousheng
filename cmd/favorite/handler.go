package main

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/favorite/pack"
	"github.com/bdyc-org/dousheng/cmd/favorite/service"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// Favorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Favorite(ctx context.Context, req *favorite.FavoriteOperationRequest) (resp *favorite.FavoriteOperationResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.FavoriteOperationResponse)

	//检查参数是否合法
	if req.UserId == 0 || req.VideoId == 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		return resp, nil
	}

	statusCode, err := service.NewFavoriteService(ctx).Favorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	switch req.ActionType {
	case 1:
		resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "点赞成功，感谢您的支持")
	case 2:
		resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "取消点赞成功")
	default:
		resp.BaseResp = pack.BuildBaseResponse(errno.ServiceErrCode, errno.ErrService.Error())
	}

	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.FavoriteListResponse)

	//检查参数是否合法
	if req.UserId == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		resp.VideoList = nil
		return resp, nil
	}

	videoList, statusCode, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.VideoList = nil
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取点赞列表成功")
	resp.VideoList = videoList
	return resp, nil

}

// FavoriteJudge implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteJudge(ctx context.Context, req *favorite.FavoriteJudgeRequest) (resp *favorite.FavoriteJudgeResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.FavoriteJudgeResponse)

	// 检查参数
	if req.UserId == 0 || len(req.VideoIds) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "判断是否点赞完成")
		resp.VideoIds = nil
		return resp, nil
	}

	videoIds, statusCode, err := service.NewFavoriteJudgeService(ctx).FavoriteJudge(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.VideoIds = nil
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "判断是否点赞完成")
	resp.VideoIds = videoIds
	return resp, nil
}

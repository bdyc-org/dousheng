package main

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/relation/pack"
	"github.com/bdyc-org/dousheng/cmd/relation/service"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Follow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Follow(ctx context.Context, req *relation.FollowRequest) (resp *relation.FollowResponse, err error) {
	resp = new(relation.FollowResponse)

	err = service.NewFollowService(ctx).Follow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.Success.WithMessage("操作成功"))
	
	return resp, nil
}

// QueryFollow implements the RelationServiceImpl interface.
// 查关注
func (s *RelationServiceImpl) QueryFollow(ctx context.Context, req *relation.QueryFollowRequest) (resp *relation.QueryFollowResponse, err error) {
	resp = new(relation.QueryFollowResponse)
	
	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
		return resp, nil
	}

	// 使用获取到的id查follow_ids
	resp.FollowIds, err = service.NewQueryFollowService(ctx).QueryFollow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.Success.WithMessage("获取用户ids成功"))

	return resp, nil
}

// QueryFollower implements the RelationServiceImpl interface.
// 查粉丝
func (s *RelationServiceImpl) QueryFollower(ctx context.Context, req *relation.QueryFollowerRequest) (resp *relation.QueryFollowerResponse, err error) {
	resp = new(relation.QueryFollowerResponse)

	if req.UserId <= 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
		return resp, nil
	}

	// 使用获取到的id查follower_ids
	resp.FollowerIds, err = service.NewQueryFollowerService(ctx).QueryFollower(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.Success.WithMessage("获取用户ids成功"))

	return resp, nil
}

// QueryUserList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) QueryUserList(ctx context.Context, req *relation.QueryUserListRequest) (resp *relation.QueryUserListResponse, err error) {
	resp = new(relation.QueryUserListResponse)
	
	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
		return resp, nil
	}

	userList, err := service.NewQueryUserListService(ctx).QueryUserList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		return resp, nil
	}
	resp.UserList = userList

	return resp, nil
}
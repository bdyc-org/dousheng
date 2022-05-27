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
	err = service.NewFollowService(ctx).Follow(req)
	if err != nil {
		return nil, err
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "操作成功")
	
	return resp, nil
}

// QueryFollow implements the RelationServiceImpl interface.
// 查关注
func (s *RelationServiceImpl) QueryFollow(ctx context.Context, req *relation.QueryFollowRequest) (resp *relation.QueryFollowResponse, err error) {
	resp = new(relation.QueryFollowResponse)
	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取用户id成功")

	// 使用获取到的id查follow_ids
	resp.FollowIds = service.NewQueryFollowService(ctx).QueryFollow(req)

	return resp, nil
}

// QueryFollower implements the RelationServiceImpl interface.
// 查粉丝
func (s *RelationServiceImpl) QueryFollower(ctx context.Context, req *relation.QueryFollowerRequest) (resp *relation.QueryFollowerResponse, err error) {
	resp = new(relation.QueryFollowerResponse)
	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取用户id成功")

	// 使用获取到的id查follower_ids
	resp.FollowerIds = service.NewQueryFollowerService(ctx).QueryFollower(req)

	return resp, nil
}

// QueryUserList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) QueryUserList(ctx context.Context, req *relation.QueryUserListRequest) (resp *relation.QueryUserListResponse, err error) {
	resp, err = service.NewQueryUserListService(ctx).QueryUserList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
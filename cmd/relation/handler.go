package main

import (
	"context"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/cmd/relation/pack"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// QueryFollow implements the RelationServiceImpl interface.
// 关注
func (s *RelationServiceImpl) QueryFollow(ctx context.Context, req *relation.QueryFollowRequest) (resp *relation.QueryFollowResponse, err error) {
	resp = new(relation.QueryFollowResponse)

	if req.UserId == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
	}

	

	return
}

// QueryFollower implements the RelationServiceImpl interface.
// 粉丝
func (s *RelationServiceImpl) QueryFollower(ctx context.Context, req *relation.QueryFollowerRequest) (resp *relation.QueryFollowerResponse, err error) {
	return
}

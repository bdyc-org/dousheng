package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/relation/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
)

type QueryUserListService struct {
	ctx context.Context
}

func NewQueryUserListService(ctx context.Context) *QueryUserListService {
	return &QueryUserListService {
		ctx: ctx,
	}
}

func (s *QueryUserListService) QueryUserList(req *relation.QueryUserListRequest) ([]*relation.User,error) {
	// 调用rpc
	userList, err := rpc.MGetUser(s.ctx, req)
	if err != nil {
		return nil, err
	}
	
	return userList, nil
}
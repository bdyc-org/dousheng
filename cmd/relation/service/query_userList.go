package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/relation/dal/db"
	"github.com/bdyc-org/dousheng/cmd/relation/pack"
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
	// 调用db
	userModels, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}

	userList := pack.UserList(userModels)
	
	return userList, nil
}
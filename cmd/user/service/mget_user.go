package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/user/dal/db"
	"github.com/bdyc-org/dousheng/cmd/user/pack"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type MGetUserService struct {
	ctx context.Context
}

func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// 获取用户（列表）信息
func (s *MGetUserService) MGetUser(req *user.MGetUserRequest) (users []*user.User, statusCode int64, err error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	//is_follow需要relation服务，暂未写
	return pack.Users(modelUsers), errno.SuccessCode, nil
}

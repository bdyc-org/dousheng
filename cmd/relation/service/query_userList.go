package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/relation/dal/db"
	"github.com/bdyc-org/dousheng/cmd/relation/pack"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type QueryUserListService struct {
	ctx context.Context
}

func NewQueryUserListService(ctx context.Context) *QueryUserListService {
	return &QueryUserListService {
		ctx: ctx,
	}
}

func (s *QueryUserListService) QueryUserList(req *relation.QueryUserListRequest) (*relation.QueryUserListResponse, error) {
	resp := new(relation.QueryUserListResponse)
	
	// 检查ids是否为空
	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		resp.UserList = nil
		return resp, nil
	}

	// 调用db
	userList, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取用户列表成功")
	resp.UserList = userList
	
	return resp, nil
}
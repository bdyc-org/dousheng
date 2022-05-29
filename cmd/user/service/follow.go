package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/user/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type FolloWService struct {
	ctx context.Context
}

func NewFollowService(ctx context.Context) *FolloWService {
	return &FolloWService{ctx: ctx}
}

func (s *FolloWService) Follow(req *user.FollowOperationRequest) (statusCode int64, err error) {
	err = db.Follow(s.ctx, req.FollowId, req.FollowerId, req.ActionType)
	if err != nil {
		return errno.ServiceErrCode, err
	}
	return errno.SuccessCode, nil
}

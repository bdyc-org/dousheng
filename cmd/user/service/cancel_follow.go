package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/user/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type CancelFolloWService struct {
	ctx context.Context
}

func NewCancelFollowService(ctx context.Context) *CancelFolloWService {
	return &CancelFolloWService{ctx: ctx}
}

func (s *CancelFolloWService) CancelFollow(req *user.CancelFollowRequest) (statusCode int64, err error) {
	err = db.CancelFollow(s.ctx, req.FollowId, req.FollowerId)
	if err != nil {
		return errno.ServiceErrCode, err
	}
	return errno.SuccessCode, nil
}

package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/user/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
)

type CancelFolloWService struct {
	ctx context.Context
}

func NewCancelFollowService(ctx context.Context) *CancelFolloWService {
	return &CancelFolloWService{ctx: ctx}
}

func (s *CancelFolloWService) CancelFollow(req *user.CancelFollowRequest) error {
	return db.CancelFollow(s.ctx, req.FollowId, req.FollowerId)
}

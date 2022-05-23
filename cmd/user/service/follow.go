package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/user/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
)

type FolloWService struct {
	ctx context.Context
}

func NewFollowService(ctx context.Context) *FolloWService {
	return &FolloWService{ctx: ctx}
}

func (s *FolloWService) Follow(req *user.FollowRequest) error {
	return db.Follow(s.ctx, req.FollowId, req.FollowerId)
}

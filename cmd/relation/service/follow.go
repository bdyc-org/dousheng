package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/relation/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
)

type FollowService struct {
	ctx context.Context
}

func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService {
		ctx: ctx,
	}
}

func (s *FollowService) Follow(req *relation.FollowRequest) error {
	var err error

	r := db.Relation{
		Follow_id: req.UserId,
		Follower_id: req.ToUserId,
	}

	// 关注
	if req.ActionType == 1 {
		err = db.Follow(s.ctx, &r)
	}
	// 取关
	if req.ActionType == 2 {
		err = db.CancelFollow(s.ctx, &r)
	}

	if err != nil {
		return err
	}
	return nil
}
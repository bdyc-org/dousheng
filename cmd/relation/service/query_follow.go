package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/relation/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
)

type QueryFollowService struct {
	ctx context.Context
}

func NewQueryFollowService(ctx context.Context) *QueryFollowService {
	return &QueryFollowService {
		ctx: ctx,
	}
}

func (s *QueryFollowService) QueryFollow(req *relation.QueryFollowRequest) []int64 {
	return db.QueryFollow(s.ctx, req.UserId)
}
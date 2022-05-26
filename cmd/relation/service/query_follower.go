package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/relation/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
)

type QueryFollowerService struct {
	ctx context.Context
}

func NewQueryFollowerService(ctx context.Context) *QueryFollowerService {
	return &QueryFollowerService {
		ctx: ctx,
	}
}

func (s *QueryFollowerService) QueryFollower(req *relation.QueryFollowerRequest) []int64 {
	return db.QueryFollower(s.ctx, req.UserId)
}
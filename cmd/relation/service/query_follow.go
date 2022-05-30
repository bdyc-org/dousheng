package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/relation/dal/db"
	"github.com/bdyc-org/dousheng/cmd/relation/pack"
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

func (s *QueryFollowService) QueryFollow(req *relation.QueryFollowRequest) ([]int64, error) {
	res, err := db.QueryFollow(s.ctx, req.UserId)
	userIds := make([]int64, len(res))

	if err != nil {
		return nil, err
	}

	rales := pack.Relas(res)
	for i := 0; i < len(rales); i++ {
		userIds[i] = rales[i].FollowId
	}

	return userIds, nil
}
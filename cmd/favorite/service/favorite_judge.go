package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/favorite/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type FavoriteJudgeService struct {
	ctx context.Context
}

func NewFavoriteJudgeService(ctx context.Context) *FavoriteJudgeService {
	return &FavoriteJudgeService{ctx: ctx}
}

func (s *FavoriteJudgeService) FavoriteJudge(req *favorite.FavoriteJudgeRequest) (videoIds []int64, statusCode int64, err error) {
	videoIds = make([]int64, 0)
	favorites, err := db.FavoriteJudge(s.ctx, req.UserId, req.VideoIds)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	if len(favorites) == 0 {
		return nil, errno.SuccessCode, nil
	}
	for _, f := range favorites {
		videoIds = append(videoIds, int64(f.VideoId))
	}
	return videoIds, errno.SuccessCode, nil
}

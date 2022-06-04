package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/favorite/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{ctx: ctx}
}

func (s *FavoriteService) Favorite(req *favorite.FavoriteOperationRequest) (statusCode int64, err error) {

	switch req.ActionType {
	case 1:
		//检查视频是否已经被点赞了
		favorites, err := db.FavoriteJudge(s.ctx, req.UserId, []int64{req.VideoId})
		if err != nil {
			return errno.ServiceErrCode, err
		}
		//如果没有，创建点赞记录
		if len(favorites) == 0 {
			err = db.CreateFavorite(s.ctx, []*db.Favorite{{
				UserId:  uint(req.UserId),
				VideoId: uint(req.VideoId),
			}})
			if err != nil {
				return errno.ServiceErrCode, err
			}
		} else {
			return errno.FavoriteErrCode, errno.ErrFavorite
		}
	case 2:
		if err := db.DeleteFavorite(s.ctx, req.UserId, req.VideoId); err != nil {
			return errno.ServiceErrCode, err
		}
	default:
		return errno.ServiceErrCode, errno.ErrWrongOperation
	}

	return errno.SuccessCode, nil
}

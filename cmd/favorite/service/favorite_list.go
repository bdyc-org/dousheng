package service

import (
	"context"
	"fmt"

	"github.com/bdyc-org/dousheng/cmd/favorite/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) FacoriteList(req *favorite.FavoriteListRequest) (videoList []*favorite.Video, statusCode int64, err error) {
	videoIds := make([]int64, 0)
	favorites, err := db.MGetFavorite(s.ctx, req.UserId)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	if len(favorites) == 0 {
		return nil, errno.SuccessCode, nil
	}
	for _, f := range favorites {
		videoIds = append(videoIds, int64(f.VideoId))
	}
	fmt.Println(videoIds)
	//等待video服务
	return nil, errno.SuccessCode, nil
}

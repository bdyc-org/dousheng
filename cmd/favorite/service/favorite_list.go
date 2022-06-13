package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/favorite/dal/db"
	"github.com/bdyc-org/dousheng/cmd/favorite/pack"
	"github.com/bdyc-org/dousheng/cmd/favorite/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) FavoriteList(req *favorite.FavoriteListRequest) (videoList []*favorite.Video, statusCode int64, err error) {
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

	//调用video服务
	videos, statusCode, err := rpc.MGetVideo(context.Background(), &video.MGetVideoRequest{
		UserId:   req.UserId,
		VideoIds: videoIds,
	})
	if err != nil {
		return nil, statusCode, err
	}

	videoList = pack.VideoList(videos)

	return videoList, errno.SuccessCode, nil
}

package service

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

type VideoFavoriteService struct {
	ctx context.Context
}

func NewVideoFavoriteService(ctx context.Context) *VideoFavoriteService {
	return &VideoFavoriteService{ctx: ctx}
}

func (v *VideoFavoriteService) VideoFavorite(req *video.DouyinVideoFavoriteRequest) error {
	return db.VideoFavorite(v.ctx, uint(req.VideoId), int(req.Action))
}

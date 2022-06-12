package service

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
)

type VideoCommentService struct {
	ctx context.Context
}

func NewVideoCommentService(ctx context.Context) *VideoCommentService {
	return &VideoCommentService{ctx: ctx}
}

func (v *VideoCommentService) VideoComment(req *video.DouyinVideoCommentRequest) error {
	return db.VideoComment(v.ctx, uint(req.VideoId), int(req.Action))
}

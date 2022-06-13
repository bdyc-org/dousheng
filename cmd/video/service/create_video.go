package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/video/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type CreateVideoService struct {
	ctx context.Context
}

func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

func (s *CreateVideoService) CreateUser(req *video.CreateVideoRequest) (statusCode int64, err error) {

	//将记录写入数据库
	err = db.CreateVideo(s.ctx, []*db.Video{{
		AuthorId:      req.AuthorId,
		PlayUrl:       req.PlayUrl,
		CoverUrl:      req.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}})
	if err != nil {
		return errno.ServiceErrCode, err
	}

	return errno.SuccessCode, nil
}

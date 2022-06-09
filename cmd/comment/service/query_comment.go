package service

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/comment/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type QueryCommentService struct {
	ctx context.Context
}

func NewQueryCommentService(ctx context.Context) *QueryCommentService {
	return &QueryCommentService{ctx: ctx}
}

func (s *QueryCommentService) QueryComment(req *comment.QueryCommentRequest) (commentList []*comment.Comment, statusCode int64, err error) {
	commentIds := make([]int64, 0)
	comments, err := db.QueryComment(s.ctx, req.VideoId)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	if len(comments) == 0 {
		return nil, errno.SuccessCode, err
	}
	for _, c := range comments {
		commentIds = append(commentIds, int64(c.Model.ID))
	}
	return nil, errno.SuccessCode, err
}

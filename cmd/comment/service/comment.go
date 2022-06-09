package service

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/comment/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

func (s *CommentService) Comment(req *comment.CommentRequest) (statusCode int64, err error) {
	c := db.Comment{
		User_id:  req.UserId,
		Video_id: req.VideoId,
		Content:  req.CommentText,
	}

	//发表评论
	if req.ActionType == 1 {
		err = db.CreatComment(s.ctx, &c)
	} else if req.ActionType == 2 {
		//删除评论
		err = db.DeleteComment(s.ctx, req.UserId, req.VideoId, req.CommentText)
	} else {
		err = errno.ParamErr
	}

	return errno.SuccessCode, nil
}

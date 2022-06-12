package service

import (
	"context"
	"time"

	"github.com/bdyc-org/dousheng/cmd/comment/dal/db"
	"github.com/bdyc-org/dousheng/cmd/comment/pack"
	"github.com/bdyc-org/dousheng/cmd/comment/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

func (s *CommentService) Comment(req *comment.CommentRequest) (*comment.Comment, error) {
	var err error
	
	c := db.Comment{
		User_id:  req.UserId,
		Video_id: req.VideoId,
		Content:  req.CommentText,
		Create_date: time.Now().Format("01-02"),
	}

	//发表评论
	if req.ActionType == 1 {
		err = db.CreatComment(s.ctx, &c)
	} else if req.ActionType == 2 {
		//删除评论
		err = db.DeleteComment(s.ctx, &c)
	}
	if err != nil {
		return nil, err
	}

	UserIds := make([]int64, 1)
	UserIds[0] = req.UserId
	Users, err := rpc.MGetUser(s.ctx, req.UserId, UserIds)
	resp := pack.Comment(&c, Users[0])
	
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}	

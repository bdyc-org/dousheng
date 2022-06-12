package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/comment/dal/db"
	"github.com/bdyc-org/dousheng/cmd/comment/pack"
	"github.com/bdyc-org/dousheng/cmd/comment/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
)

type QueryCommentService struct {
	ctx context.Context
}

func NewQueryCommentService(ctx context.Context) *QueryCommentService {
	return &QueryCommentService{ctx: ctx}
}

func (s *QueryCommentService) QueryComment(req *comment.QueryCommentRequest) (commentList []*comment.Comment, err error) {
	userIds := make([]int64, 0)
	comments, err := db.QueryComment(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	for _, c := range comments {
		userIds = append(userIds, c.User_id)
	}

	Users, err := rpc.MGetUser(s.ctx, req.UserId, userIds)
	if err != nil {
		return nil, err
	}

	resp := pack.Comments(comments, Users)
	return resp, nil
}

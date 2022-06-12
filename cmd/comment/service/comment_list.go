package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/comment/dal/db"
	"github.com/bdyc-org/dousheng/cmd/comment/pack"
	"github.com/bdyc-org/dousheng/cmd/comment/rpc"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (s *CommentListService) CommentList(req *comment.CommentListRequest) (commentList []*comment.Comment, statusCode int64, err error) {
	userIds := make([]int64, 0)
	comments, err := db.MGetComment(s.ctx, req.VideoId)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	for _, c := range comments {
		userIds = append(userIds, c.UserId)
	}

	users, statusCode, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{
		UserId:  req.UserId,
		UserIds: userIds,
	})
	if err != nil {
		return nil, statusCode, err
	}

	userList := pack.UserList(users)

	resp := pack.Comments(comments, userList)
	return resp, errno.SuccessCode, nil
}

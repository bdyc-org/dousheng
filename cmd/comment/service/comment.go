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

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

func (s *CommentService) Comment(req *comment.CommentRequest) (*comment.Comment, int64, error) {
	switch req.ActionType {
	case 1:
		// 创建对象记录
		item := db.Comment{
			UserId:  req.UserId,
			VideoId: req.VideoId,
			Content: *req.CommentText,
		}
		// 将评论插入comments表
		err := db.CreateComment(s.ctx, &item)
		if err != nil {
			return nil, errno.ServiceErrCode, err
		}
		// 获取评论用户信息
		users, statusCode, err := rpc.MGetUser(context.Background(), &user.MGetUserRequest{
			UserId:  req.UserId,
			UserIds: []int64{req.UserId},
		})
		if err != nil {
			return nil, statusCode, err
		}
		if len(users) == 0 {
			return nil, errno.ServiceErrCode, errno.ErrService
		}
		user := pack.User(users[0])
		// 返回评论内容
		comment := pack.Comment(&item, user)
		return comment, errno.SuccessCode, nil
	case 2:
		err := db.DeleteComment(s.ctx, *req.CommentId)
		if err != nil {
			return nil, errno.ServiceErrCode, err
		}
		return nil, errno.SuccessCode, nil
	default:
		return nil, errno.ServiceErrCode, errno.ErrService
	}
}

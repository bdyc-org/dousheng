package main

import (
	"context"
	"github.com/bdyc-org/dousheng/cmd/comment/pack"
	"github.com/bdyc-org/dousheng/cmd/comment/service"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// Comment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Comment(ctx context.Context, req *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.CommentResponse)

	//检查参数是否合法
	if req.UserId == 0 || req.VideoId == 0 || req.CommentText == "" || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		return resp, nil
	}

	comment, err := service.NewCommentService(ctx).Comment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(errno.ServiceErrCode, err.Error())
		return resp, nil
	}
	resp.Comment = comment

	switch req.ActionType {
	case 1:
		resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "评论成功")
	case 2:
		resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "删除评论成功")
	default:
		resp.BaseResp = pack.BuildBaseResponse(errno.ServiceErrCode, errno.ErrService.Error())
	}

	return resp, nil
}

// QueryComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) QueryComment(ctx context.Context, req *comment.QueryCommentRequest) (resp *comment.QueryCommentResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.QueryCommentResponse)

	if req.VideoId == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		resp.CommentList = nil
		return resp, nil
	}

	commentList, err := service.NewQueryCommentService(ctx).QueryComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(errno.ServiceErrCode, err.Error())
		return resp, err
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取评论列表成功")
	resp.CommentList = commentList
	return resp, nil
}

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
	resp.Comment = nil

	//检查参数是否合法
	if req.UserId == 0 || req.VideoId == 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		return resp, nil
	}
	if (req.ActionType == 1 && req.CommentText == nil) || (req.ActionType == 2 && req.CommentId == nil) {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		return resp, nil
	}

	comment, statusCode, err := service.NewCommentService(ctx).Comment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	switch req.ActionType {
	case 1:
		resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "评论成功")
		resp.Comment = comment
	case 2:
		resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "删除评论成功")
	default:
		resp.BaseResp = pack.BuildBaseResponse(errno.ServiceErrCode, errno.ErrService.Error())
	}

	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.CommentListResponse)
	resp.CommentList = nil

	if req.VideoId == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		return resp, nil
	}

	commentList, statusCode, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取评论列表成功")
	resp.CommentList = commentList
	return resp, nil
}

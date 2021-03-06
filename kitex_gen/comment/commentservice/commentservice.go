// Code generated by Kitex v0.3.1. DO NOT EDIT.

package commentservice

import (
	"context"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Comment":     kitex.NewMethodInfo(commentHandler, newCommentServiceCommentArgs, newCommentServiceCommentResult, false),
		"CommentList": kitex.NewMethodInfo(commentListHandler, newCommentServiceCommentListArgs, newCommentServiceCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func commentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCommentArgs)
	realResult := result.(*comment.CommentServiceCommentResult)
	success, err := handler.(comment.CommentService).Comment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentArgs() interface{} {
	return comment.NewCommentServiceCommentArgs()
}

func newCommentServiceCommentResult() interface{} {
	return comment.NewCommentServiceCommentResult()
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentServiceCommentListArgs)
	realResult := result.(*comment.CommentServiceCommentListResult)
	success, err := handler.(comment.CommentService).CommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCommentListArgs() interface{} {
	return comment.NewCommentServiceCommentListArgs()
}

func newCommentServiceCommentListResult() interface{} {
	return comment.NewCommentServiceCommentListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Comment(ctx context.Context, req *comment.CommentRequest) (r *comment.CommentResponse, err error) {
	var _args comment.CommentServiceCommentArgs
	_args.Req = req
	var _result comment.CommentServiceCommentResult
	if err = p.c.Call(ctx, "Comment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, req *comment.CommentListRequest) (r *comment.CommentListResponse, err error) {
	var _args comment.CommentServiceCommentListArgs
	_args.Req = req
	var _result comment.CommentServiceCommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

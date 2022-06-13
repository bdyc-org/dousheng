package rpc

import (
	"context"
	"errors"
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/bdyc-org/dousheng/kitex_gen/comment/commentservice"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/bdyc-org/dousheng/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var commentClient commentservice.Client

func initCommentRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		constants.CommentServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

func CommentOperation(ctx context.Context, req *comment.CommentRequest) (comment *comment.Comment, statusCode int64, err error) {
	resp, err := commentClient.Comment(ctx, req)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.Comment, errno.SuccessCode, nil
}

func CommentList(ctx context.Context, req *comment.CommentListRequest) (commentList []*comment.Comment, statusCode int64, err error) {
	resp, err := commentClient.CommentList(ctx, req)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.CommentList, errno.SuccessCode, nil
}

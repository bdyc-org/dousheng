package rpc

import (
	"context"
	"time"

	"github.com/bdyc-org/dousheng/cmd/relation/pack"
	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/kitex_gen/relation/relationservice"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/bdyc-org/dousheng/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelationRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := relationservice.NewClient(
		constants.RelationServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

func RelaFollow(ctx context.Context, req *relation.FollowRequest) (resp *relation.FollowResponse, err error) {
	resp = new(relation.FollowResponse)

	// 调用userClient
	res, err := UserFollow(ctx, &user.FollowOperationRequest{
		FollowId: req.UserId,
		FollowerId: req.ToUserId,
		ActionType: req.ActionType,
	})

	if res != errno.SuccessCode {
		resp.BaseResp = pack.BuildBaseResponse(errno.NewErrNo(res, err.Error()))
		return resp, err
	}

	// 调用relationClient的Follow
	return relationClient.Follow(ctx, req)
}

func QueryUserList(ctx context.Context, req *relation.QueryUserListRequest) (resp *relation.QueryUserListResponse, err error) {
	return relationClient.QueryUserList(ctx, req)
}

func QueryFollow(ctx context.Context, userId int64) (resp *relation.QueryFollowResponse, err error) {
	return relationClient.QueryFollow(ctx, &relation.QueryFollowRequest{UserId: userId})
}

func QueryFollower(ctx context.Context, userId int64) (resp *relation.QueryFollowerResponse, err error) {
	return relationClient.QueryFollower(ctx, &relation.QueryFollowerRequest{UserId: userId})
}
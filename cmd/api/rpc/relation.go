package rpc

import (
	"context"
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/relation"
	"github.com/bdyc-org/dousheng/kitex_gen/relation/relationservice"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/constants"
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
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
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
	// 调用userClient的Follow
	_, err = userClient.Follow(ctx, &user.FollowRequest{
		FollowId: req.UserId,
		FollowerId: req.ToUserId,
	})
	if err != nil {
		return nil, err
	}
	
	// 调用relationClient的Follow
	resp, err = relationClient.Follow(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func QueryUserList(ctx context.Context, req *relation.QueryUserListRequest) (resp *relation.QueryUserListResponse, err error) {
	resp, err = relationClient.QueryUserList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func QueryFollow(ctx context.Context, userId int64) (resp *relation.QueryFollowResponse, err error) {
	resp, err = relationClient.QueryFollow(ctx, &relation.QueryFollowRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func QueryFollower(ctx context.Context, userId int64) (resp *relation.QueryFollowerResponse, err error) {
	resp, err = relationClient.QueryFollower(ctx, &relation.QueryFollowerRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
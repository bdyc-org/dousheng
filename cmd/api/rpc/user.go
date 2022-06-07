package rpc

import (
	"context"
	"errors"
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/kitex_gen/user/userservice"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/bdyc-org/dousheng/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
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
	userClient = c
}

func CreateUser(ctx context.Context, req *user.CreateUserRequest) (user_id int64, statusCode int64, err error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return 0, errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.UserId, errno.SuccessCode, nil
}

func CheckUser(ctx context.Context, req *user.CheckUserRequest) (user_id int64, statusCode int64, err error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.UserId, errno.SuccessCode, nil
}

func MGetUser(ctx context.Context, req *user.MGetUserRequest) (user []*user.User, statusCode int64, err error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.UserList, errno.SuccessCode, nil
}

func UserFollow(ctx context.Context, req *user.FollowOperationRequest) (resp *user.FollowOperationResponse, err error) {
	return userClient.Follow(ctx, req)
}

func Authentication(ctx context.Context, req *user.AuthenticationRequest) (user_id int64, statusCode int64, err error) {
	resp, err := userClient.Authentication(ctx, req)
	if err != nil {
		return 0, errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.UserId, errno.SuccessCode, nil
}

func UserFavorite(ctx context.Context, req *user.FavoriteOperationRequest) (statusCode int64, err error) {
	resp, err := userClient.Favorite(ctx, req)
	if err != nil {
		return errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return errno.SuccessCode, nil
}

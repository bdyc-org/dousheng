package rpc

import (
	"context"
	"errors"
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/kitex_gen/video/videoservice"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/bdyc-org/dousheng/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)

	if err != nil {
		panic(err)
	}
	videoClient = c
}

func MGetVideo(ctx context.Context, req *video.MGetVideoRequest) (videoList []*video.Video, statusCode int64, err error) {
	resp, err := videoClient.MGetVideo(ctx, req)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, errno.SuccessCode, nil
}

func VideoFavorite(ctx context.Context, req *video.FavoriteOperationRequest) (statusCode int64, err error) {
	resp, err := videoClient.Favorite(ctx, req)
	if err != nil {
		return errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return errno.SuccessCode, nil
}

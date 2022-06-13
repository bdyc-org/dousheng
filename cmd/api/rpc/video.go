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

func CreateVideo(ctx context.Context, req *video.CreateVideoRequest) (statusCode int64, err error) {
	resp, err := videoClient.CreateVideo(ctx, req)
	if err != nil {
		return errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return errno.SuccessCode, nil
}

func Feed(ctx context.Context, req *video.FeedRequest) (nextTime int64, videoList []*video.Video, statusCode int64, err error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return 0, nil, errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, nil, resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.NextTime, resp.VideoList, errno.SuccessCode, nil
}

func PublishList(ctx context.Context, req *video.PublishListRequest) (videoList []*video.Video, statusCode int64, err error) {
	resp, err := videoClient.PublishList(ctx, req)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, errno.SuccessCode, nil
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

func VideoComment(ctx context.Context, req *video.CommentOperationRequest) (statusCode int64, err error) {
	resp, err := videoClient.Comment(ctx, req)
	if err != nil {
		return errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	return errno.SuccessCode, nil
}

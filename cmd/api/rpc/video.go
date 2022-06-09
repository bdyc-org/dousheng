package rpc

import (
	"context"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/bdyc-org/dousheng/kitex_gen/video/videoservice"
	"github.com/bdyc-org/dousheng/pkg/constants"
	error2 "github.com/bdyc-org/dousheng/pkg/error"
	"github.com/bdyc-org/dousheng/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
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
		client.WithMuxConnection(1),          // mux
		client.WithRPCTimeout(3*time.Second), //rpc timeout
		//client.WithConnectTimeout()
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	videoClient = c
}

func PublicVideo(ctx context.Context, req *video.DouyinPublishActionRequest) error {
	resp, err := videoClient.PublishAction(ctx, req)
	if err != nil {
		return err
	}
	// TODO StatusCode need  change
	if resp.StatusCode != 0 {
		return error2.NewErrNo(resp.StatusCode, *(resp.StatusMsg))
	}
	return nil
}
func FeedVideo(ctx context.Context, req *video.DouyinFeedRequest) ([]*video.Video, *int64, error) {
	resp, err := videoClient.FeedVideo(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != 0 {
		return nil, nil, error2.NewErrNo(resp.StatusCode, *(resp.StatusMsg))
	}

	return resp.VideoList, resp.NextTime, err
}

func PublishList(ctx context.Context, req *video.DouyinPublishListRequest) ([]*video.Video, error) {
	resp, err := videoClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, error2.NewErrNo(resp.StatusCode, *(resp.StatusMsg))
	}

	return resp.VideoList, nil
}

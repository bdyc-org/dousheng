package rpc

import (
	"context"
	favorite2 "github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite/favoriteservice"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/bdyc-org/dousheng/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		constants.UserServiceName,
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
	favoriteClient = c
}

func FavoriteJudge(ctx context.Context, req *favorite2.FavoriteJudgeRequest) (videoIds []int64, err error) {

	videoIds = make([]int64, 0)
	resp, err := favoriteClient.FavoriteJudge(ctx, req)
	if err != nil {
		panic(err)
	}

	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	videoIds = resp.VideoIds
	return videoIds, nil
}

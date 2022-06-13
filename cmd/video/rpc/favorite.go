package rpc

import (
	"context"
	"errors"
	"time"

	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite/favoriteservice"
	"github.com/bdyc-org/dousheng/pkg/constants"
	"github.com/bdyc-org/dousheng/pkg/errno"
	"github.com/bdyc-org/dousheng/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := favoriteservice.NewClient(
		constants.FavoriteServiceName,
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

func FavoriteJudge(ctx context.Context, req *favorite.FavoriteJudgeRequest) (videoIds map[int64]bool, statusCode int64, err error) {
	resp, err := favoriteClient.FavoriteJudge(ctx, req)
	if err != nil {
		return nil, errno.ServiceErrCode, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, resp.BaseResp.StatusCode, errors.New(resp.BaseResp.StatusMsg)
	}
	res := make(map[int64]bool)
	for _, videoID := range resp.VideoIds {
		res[videoID] = true
	}
	return res, errno.SuccessCode, nil
}

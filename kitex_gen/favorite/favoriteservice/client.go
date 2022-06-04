// Code generated by Kitex v0.3.1. DO NOT EDIT.

package favoriteservice

import (
	"context"
	"github.com/bdyc-org/dousheng/kitex_gen/favorite"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Favorite(ctx context.Context, req *favorite.FavoriteOperationRequest, callOptions ...callopt.Option) (r *favorite.FavoriteOperationResponse, err error)
	FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest, callOptions ...callopt.Option) (r *favorite.FavoriteListResponse, err error)
	FavoriteJudge(ctx context.Context, req *favorite.FavoriteJudgeRequest, callOptions ...callopt.Option) (r *favorite.FavoriteJudgeResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFavoriteServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFavoriteServiceClient struct {
	*kClient
}

func (p *kFavoriteServiceClient) Favorite(ctx context.Context, req *favorite.FavoriteOperationRequest, callOptions ...callopt.Option) (r *favorite.FavoriteOperationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Favorite(ctx, req)
}

func (p *kFavoriteServiceClient) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest, callOptions ...callopt.Option) (r *favorite.FavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, req)
}

func (p *kFavoriteServiceClient) FavoriteJudge(ctx context.Context, req *favorite.FavoriteJudgeRequest, callOptions ...callopt.Option) (r *favorite.FavoriteJudgeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteJudge(ctx, req)
}
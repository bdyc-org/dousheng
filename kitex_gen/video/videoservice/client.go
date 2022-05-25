// Code generated by Kitex v0.3.1. DO NOT EDIT.

package videoservice

import (
	"context"
	"github.com/bdyc-org/dousheng/kitex_gen/video"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FeedVideo(ctx context.Context, req *video.DouyinFeedRequest, callOptions ...callopt.Option) (r *video.DouyinFeedResponse, err error)
	PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *video.DouyinPublishActionResponse, err error)
	PublishList(ctx context.Context, req *video.DouyinPublishListRequest, callOptions ...callopt.Option) (r *video.DouyinPublishListResponse, err error)
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
	return &kVideoServiceClient{
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

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) FeedVideo(ctx context.Context, req *video.DouyinFeedRequest, callOptions ...callopt.Option) (r *video.DouyinFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FeedVideo(ctx, req)
}

func (p *kVideoServiceClient) PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *video.DouyinPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishAction(ctx, req)
}

func (p *kVideoServiceClient) PublishList(ctx context.Context, req *video.DouyinPublishListRequest, callOptions ...callopt.Option) (r *video.DouyinPublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, req)
}

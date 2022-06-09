// Code generated by Kitex v0.3.1. DO NOT EDIT.

package commentservice

import (
	"context"
	"github.com/bdyc-org/dousheng/kitex_gen/comment"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Comment(ctx context.Context, req *comment.CommentRequest, callOptions ...callopt.Option) (r *comment.CommentResponse, err error)
	QueryComment(ctx context.Context, req *comment.QueryCommentRequest, callOptions ...callopt.Option) (r *comment.QueryCommentResponse, err error)
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
	return &kCommentServiceClient{
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

type kCommentServiceClient struct {
	*kClient
}

func (p *kCommentServiceClient) Comment(ctx context.Context, req *comment.CommentRequest, callOptions ...callopt.Option) (r *comment.CommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Comment(ctx, req)
}

func (p *kCommentServiceClient) QueryComment(ctx context.Context, req *comment.QueryCommentRequest, callOptions ...callopt.Option) (r *comment.QueryCommentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryComment(ctx, req)
}
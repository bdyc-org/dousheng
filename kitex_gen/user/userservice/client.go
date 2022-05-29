// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"context"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest, callOptions ...callopt.Option) (r *user.CreateUserResponse, err error)
	CheckUser(ctx context.Context, req *user.CheckUserRequest, callOptions ...callopt.Option) (r *user.CheckUserResponse, err error)
	MGetUser(ctx context.Context, req *user.MGetUserRequest, callOptions ...callopt.Option) (r *user.MGetUserResponse, err error)
	Follow(ctx context.Context, req *user.FollowOperationRequest, callOptions ...callopt.Option) (r *user.FollowOperationResponse, err error)
	Authentication(ctx context.Context, req *user.AuthenticationRequest, callOptions ...callopt.Option) (r *user.AuthenticationResponse, err error)
	Favorite(ctx context.Context, req *user.FavoriteOperationRequest, callOptions ...callopt.Option) (r *user.FavoriteOperationResponse, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, req *user.CreateUserRequest, callOptions ...callopt.Option) (r *user.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, req)
}

func (p *kUserServiceClient) CheckUser(ctx context.Context, req *user.CheckUserRequest, callOptions ...callopt.Option) (r *user.CheckUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckUser(ctx, req)
}

func (p *kUserServiceClient) MGetUser(ctx context.Context, req *user.MGetUserRequest, callOptions ...callopt.Option) (r *user.MGetUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetUser(ctx, req)
}

func (p *kUserServiceClient) Follow(ctx context.Context, req *user.FollowOperationRequest, callOptions ...callopt.Option) (r *user.FollowOperationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Follow(ctx, req)
}

func (p *kUserServiceClient) Authentication(ctx context.Context, req *user.AuthenticationRequest, callOptions ...callopt.Option) (r *user.AuthenticationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Authentication(ctx, req)
}

func (p *kUserServiceClient) Favorite(ctx context.Context, req *user.FavoriteOperationRequest, callOptions ...callopt.Option) (r *user.FavoriteOperationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Favorite(ctx, req)
}
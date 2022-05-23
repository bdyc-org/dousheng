package main

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/user/pack"
	"github.com/bdyc-org/dousheng/cmd/user/service"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CreateUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
		return resp, nil
	}

	user_id, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		resp.UserId = 0
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.Success)
	resp.UserId = user_id
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CheckUserResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
		return resp, nil
	}

	user_id, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		resp.UserId = 0
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.Success)
	resp.UserId = user_id
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.MGetUserResponse)

	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		resp.UserList = nil
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.Success)
	resp.UserList = users
	return resp, nil
}

// NewFollow_ implements the UserServiceImpl interface.
func (s *UserServiceImpl) Follow(ctx context.Context, req *user.FollowRequest) (resp *user.FollowResponse, err error) {
	// TODO: Your code here...
	resp = new(user.FollowResponse)

	if req.FollowId == 0 || req.FollowerId == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
	}

	err = service.NewFollowService(ctx).Follow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.Success)
	return resp, nil
}

// CancelFollow implements the UserServiceImpl interface.
func (s *UserServiceImpl) CancelFollow(ctx context.Context, req *user.CancelFollowRequest) (resp *user.CancelFollowResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CancelFollowResponse)

	if req.FollowId == 0 || req.FollowerId == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
	}

	err = service.NewCancelFollowService(ctx).CancelFollow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.Success)
	return resp, nil
}

func (s *UserServiceImpl) Authentication(ctx context.Context, req *user.AuthenticationRequest) (resp *user.AuthenticationResponse, err error) {
	// TODO: Your code here...
	resp = new(user.AuthenticationResponse)

	if len(req.Username) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErr)
		return resp, nil
	}

	user_id, err := service.NewAuthenticationService(ctx).Authentication(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(err)
		resp.UserId = 0
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResponse(errno.Success)
	resp.UserId = user_id
	return resp, nil
}

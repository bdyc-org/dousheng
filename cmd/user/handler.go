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

	//检查参数是否合法
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		resp.UserId = 0
		return resp, nil
	}

	//将用户名密码插入数据库，返回user_id
	user_id, statusCode, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.UserId = 0
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "注册用户成功")
	resp.UserId = user_id
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CheckUserResponse)

	//检查参数是否合法
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		resp.UserId = 0
		return resp, nil
	}

	//查询数据库，看用户名和密码是否正确，正确返回user_id
	user_id, statusCode, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.UserId = 0
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "用户名，密码正确")
	resp.UserId = user_id
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.MGetUserResponse)

	//检查参数是否合法
	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
		resp.UserList = nil
		return resp, nil
	}

	users, statusCode, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.UserList = nil
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "获取用户信息成功")
	resp.UserList = users
	return resp, nil
}

// NewFollow_ implements the UserServiceImpl interface.
func (s *UserServiceImpl) Follow(ctx context.Context, req *user.FollowRequest) (resp *user.FollowResponse, err error) {
	// TODO: Your code here...
	resp = new(user.FollowResponse)

	//检查参数是否合法
	if req.FollowId == 0 || req.FollowerId == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
	}

	statusCode, err := service.NewFollowService(ctx).Follow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "关注成功")
	return resp, nil
}

// CancelFollow implements the UserServiceImpl interface.
func (s *UserServiceImpl) CancelFollow(ctx context.Context, req *user.CancelFollowRequest) (resp *user.CancelFollowResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CancelFollowResponse)

	//检查参数是否合法
	if req.FollowId == 0 || req.FollowerId == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.ParamErrCode, errno.Errparameter.Error())
	}

	statusCode, err := service.NewCancelFollowService(ctx).CancelFollow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "取消关注成功")
	return resp, nil
}

func (s *UserServiceImpl) Authentication(ctx context.Context, req *user.AuthenticationRequest) (resp *user.AuthenticationResponse, err error) {
	// TODO: Your code here...
	resp = new(user.AuthenticationResponse)

	//用户名为空，可能是用户为登录，Token过期或不可用，或者Token解析失败
	if len(req.Username) == 0 {
		resp.BaseResp = pack.BuildBaseResponse(errno.TokenInvalidErrCode, errno.ErrTokenInvalid.Error())
		resp.UserId = 0
		return resp, nil
	}

	//通过Token解析出来的username查找并返回user_id
	user_id, statusCode, err := service.NewAuthenticationService(ctx).Authentication(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResponse(statusCode, err.Error())
		resp.UserId = 0
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResponse(errno.SuccessCode, "token鉴权成功")
	resp.UserId = user_id
	return resp, nil
}

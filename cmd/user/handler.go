package main

import (
	"context"

	"github.com/bdyc-org/dousheng/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// NewFollow_ implements the UserServiceImpl interface.
func (s *UserServiceImpl) NewFollow_(ctx context.Context, req *user.NewFollowerRequest_) (resp *user.NewFollowResponse_, err error) {
	// TODO: Your code here...
	return
}

// NewFollower_ implements the UserServiceImpl interface.
func (s *UserServiceImpl) NewFollower_(ctx context.Context, req *user.NewFollowerRequest_) (resp *user.NewFollowerRequest_, err error) {
	// TODO: Your code here...
	return
}

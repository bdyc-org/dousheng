package service

import (
	"context"

	"github.com/bdyc-org/dousheng/cmd/user/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type AuthenticationService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewAuthenticationService(ctx context.Context) *AuthenticationService {
	return &AuthenticationService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *AuthenticationService) Authentication(req *user.AuthenticationRequest) (int64, error) {
	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	return int64(u.ID), nil
}

package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/bdyc-org/dousheng/cmd/user/dal/db"
	"github.com/bdyc-org/dousheng/kitex_gen/user"
	"github.com/bdyc-org/dousheng/pkg/errno"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) (int64, error) {
	fmt.Println("我来了，user服务的service层")
	//判断用户是否存在
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		fmt.Println("我来了，db.QueryUser出错了")
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	//对密码进行加密
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	//将记录写入数据库
	err = db.CreateUser(s.ctx, []*db.User{{
		Name:          req.Username,
		Password:      passWord,
		FollowCount:   0,
		FollowerCount: 0,
	}})
	if err != nil {
		return 0, err
	}

	//在数据库查询user_id并返回
	users, err = db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}

	return int64(users[0].ID), nil
}

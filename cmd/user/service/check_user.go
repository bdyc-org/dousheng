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

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// 检查用户名和密码是否正确，正确返回user_id
func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (user_id int64, statusCode int64, err error) {
	//对密码进行加密
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, errno.ServiceErrCode, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	//通过用户名查询
	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, errno.ServiceErrCode, err
	}

	//如果结果为空，则表明用户不存在
	if len(users) == 0 {
		return 0, errno.UserNotExistErrCode, errno.ErrUserNotExist
	}

	//如果密码不匹配，则登录失败
	u := users[0]
	if u.Password != passWord {
		return 0, errno.LoginErrCode, errno.ErrLogin
	}
	return int64(u.ID), errno.SuccessCode, nil
}

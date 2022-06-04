package errno

import (
	"errors"
)

const (
	SuccessCode            int64 = 0
	ServiceErrCode         int64 = 10001
	ParamErrCode           int64 = 10002
	LoginErrCode           int64 = 10003
	UserNotExistErrCode    int64 = 10004
	UserNameHasUsedErrCode int64 = 10005
	TokenInvalidErrCode    int64 = 10006
	FavoriteErrCode        int64 = 10007
)

var (
	ErrService         error = errors.New("服务异常，请稍后再试")
	Errparameter       error = errors.New("参数不正确")
	ErrLogin           error = errors.New("用户名或密码错误")
	ErrUserNotExist    error = errors.New("用户不存在")
	ErrUserNameHasUsed error = errors.New("用户名已经被使用")
	ErrTokenInvalid    error = errors.New("token已过期或不可用，请重新登录")
	ErrWrongOperation  error = errors.New("系统错误或操作不合法")
	ErrFavorite        error = errors.New("该视频已经被您点赞了")
)
